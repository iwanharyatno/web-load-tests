<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Participant;
use App\Models\Payment;
use App\Models\Ticket;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class RegisterController extends Controller
{
    public function __invoke(Request $request): JsonResponse
    {
        $validated = $request->validate([
            'name' => 'required|string|max:255',
            'email' => 'required|email|unique:participants,email',
            'phone' => 'required|string|max:50',
            'ticket_id' => 'required|exists:tickets,id',
        ]);

        return DB::transaction(function () use ($validated) {
            $participant = Participant::create([
                'name' => $validated['name'],
                'email' => $validated['email'],
                'phone' => $validated['phone'],
            ]);

            $ticket = Ticket::find($validated['ticket_id']);

            $orderId = 'REG-' . $participant->id . '-' . now()->timestamp;

            $payment = Payment::create([
                'participant_id' => $participant->id,
                'ticket_id' => $ticket->id,
                'order_id' => $orderId,
                'subtotal' => $ticket->price,
                'status' => 'pending',
            ]);

            return response()->json([
                'success' => true,
                'data' => [
                    'participant' => $participant,
                    'payment' => $payment,
                ],
            ], 201);
        });
    }
}
