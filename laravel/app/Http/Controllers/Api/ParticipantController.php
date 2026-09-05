<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Participant;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class ParticipantController extends Controller
{
    public function show(int $id): JsonResponse
    {
        $participant = Participant::find($id);

        if (!$participant) {
            return response()->json([
                'success' => false,
                'error' => 'participant not found',
            ], 404);
        }

        $payments = $participant->payments;

        return response()->json([
            'success' => true,
            'data' => [
                'participant' => $participant,
                'payments' => $payments,
            ],
        ]);
    }
}
