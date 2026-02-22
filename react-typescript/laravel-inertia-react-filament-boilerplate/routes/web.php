<?php

use Illuminate\Support\Facades\Route;
use Inertia\Inertia;
use Nette\Utils\Json;

Route::get('/', function () {
    return Inertia::render('Landing');
});

Route::get('/test', function () {
    return response()->json(['message' => 'Hello World']);
});