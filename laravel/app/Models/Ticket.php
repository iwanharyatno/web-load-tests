<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Ticket extends Model
{
    use HasFactory;

    protected $fillable = ['name', 'price', 'bib_prefix', 'bib_padding', 'bib_increment'];

    public function payments(): HasMany
    {
        return $this->hasMany(Payment::class);
    }
}
