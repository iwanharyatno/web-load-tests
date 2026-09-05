<?php

namespace App\Services;

class BibService
{
    public static function generateBibNumber(string $prefix, int $padding, int $increment): string
    {
        return $prefix . str_pad((string) $increment, $padding, '0', STR_PAD_LEFT);
    }
}
