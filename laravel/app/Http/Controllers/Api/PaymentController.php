<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Payment;
use Illuminate\Http\JsonResponse;

class PaymentController extends Controller
{
    public function show(string $orderId): JsonResponse
    {
        $payment = Payment::where('order_id', $orderId)->first();

        if (!$payment) {
            return response()->json([
                'success' => false,
                'error' => 'payment not found',
            ], 404);
        }

        return response()->json([
            'success' => true,
            'data' => $payment,
        ]);
    }
}
