<?php
namespace vendor\laravel\framework\src;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class unique extends Model
{
    //use HasFactory;
    public function first(){
    $a = DB::select("select unique_id from unique_num");
    echo a;
    }
}
