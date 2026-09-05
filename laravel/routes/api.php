<?php

use App\Http\Controllers\Api\ParticipantController;
use App\Http\Controllers\Api\PaymentController;
use App\Http\Controllers\Api\RegisterController;
use App\Http\Controllers\Api\TicketController;
use App\Http\Controllers\Api\WebhookController;
use Illuminate\Support\Facades\Route;

Route::get('/tickets', [TicketController::class, 'index']);
Route::post('/register', RegisterController::class);
Route::get('/participants/{id}', [ParticipantController::class, 'show']);
Route::get('/payments/{orderId}', [PaymentController::class, 'show']);
Route::post('/webhook/payment', WebhookController::class);
