<?php

namespace Database\Seeders;

use App\Models\Ticket;
use Illuminate\Database\Seeder;

class TicketSeeder extends Seeder
{
    public function run(): void
    {
        $tickets = [
            ['name' => 'Early Bird', 'price' => 25.00, 'bib_prefix' => 'EB', 'bib_padding' => 5, 'bib_increment' => 1],
            ['name' => 'Regular', 'price' => 50.00, 'bib_prefix' => 'RG', 'bib_padding' => 5, 'bib_increment' => 1],
            ['name' => 'VIP', 'price' => 100.00, 'bib_prefix' => 'VIP', 'bib_padding' => 4, 'bib_increment' => 1],
        ];

        foreach ($tickets as $ticket) {
            Ticket::create($ticket);
        }
    }
}
