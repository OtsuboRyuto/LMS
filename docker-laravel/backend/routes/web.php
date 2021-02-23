<?php

use Illuminate\Support\Facades\Route;
/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::match(['get', 'post'], 'bbs/{hash}/{id}', "App\Http\Controllers\FirstController@getId");
Route::get("bbs/id/{id}/group/{group}/content/{content}", "App\Http\Controllers\FirstController@postContent");
Route::get("bbs/group/{group}/flag/{false}", "App\Http\Controllers\FirstController@getContent");
Route::get('bbs/a', function () {
    //return redirect('bbs', "App\Http\Controllers\FirstController@getId");
    return "asd";
});