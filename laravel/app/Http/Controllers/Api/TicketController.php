<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\Ticket;
use Illuminate\Http\JsonResponse;

class TicketController extends Controller
{
    public function index(): JsonResponse
    {
        $tickets = Ticket::all();

        return response()->json([
            'success' => true,
            'data' => $tickets,
        ]);
    }
}
