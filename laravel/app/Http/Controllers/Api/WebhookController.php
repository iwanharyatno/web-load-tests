<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Payment;
use App\Models\Ticket;
use App\Services\BibService;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class WebhookController extends Controller
{
    public function __invoke(Request $request): JsonResponse
    {
        $validated = $request->validate([
            'order_id' => 'required|string',
        ]);

        $payment = Payment::where('order_id', $validated['order_id'])->first();

        if (!$payment) {
            return response()->json([
                'success' => false,
                'error' => 'payment not found',
            ], 404);
        }

        if ($payment->status === 'paid') {
            return response()->json([
                'success' => true,
                'message' => 'payment already processed',
                'data' => $payment,
            ]);
        }

        return DB::transaction(function () use ($payment) {
            $ticket = Ticket::find($payment->ticket_id);

            $bibNumber = BibService::generateBibNumber(
                $ticket->bib_prefix,
                $ticket->bib_padding,
                $ticket->bib_increment
            );

            $ticket->increment('bib_increment');

            $payment->update([
                'status' => 'paid',
                'bib_number' => $bibNumber,
            ]);

            return response()->json([
                'success' => true,
                'message' => 'payment confirmed',
                'data' => $payment->fresh(),
            ]);
        });
    }
}
