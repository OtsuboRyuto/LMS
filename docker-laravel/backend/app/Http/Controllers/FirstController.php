<?php
//モデルの雛形が作成できないため、原因が特定できるまではコントローラーで代用する
namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\unique;
use Illuminate\Support\Facades\DB;

class FirstController extends Controller
{
	public function getContent($group, $flag){//jsonのajax受付としても利用するため、flagによって見極める
   	$data = $contents = DB::table("bbs")->where("group_id", $group)->get();
   	$contents_id = $contents -> pluck("id");
   	$content = $contents -> pluck("contents");
   	//echo response()->json(["contents" => $content, "contents_id", $contents_id]);
   	if ($flag == true)
   		return ["contents_id" => $contents_id, "contents" => $content];
   	else
   		return response()->json(["contents" => $content, "contents_id", $contents_id]);
   }
	public function getId($hash_val, $id){
		//urlのクエリのハッシュIDより、IDを取り出し、そこからユニークIDを取り出す。
		$groups = DB::table("how_many")->get()->pluck("num_of_groups");
		$dehash = dechex($hash_val);//10進数を16進数に変換
		$id_indb = DB::table("auth")->where("user_print", $dehash)->get()->pluck("id")[0];//このハッシュは一度ハッシュ値(16進数)にしたものを8桁に切り,それを10進数にしたものであるため、16進数に戻して検索する
		if ($id!=$id_indb)//認証チェック
			return "値が不正です";
   		$unique = DB::table("unique_num")->where("id", $id_indb)->get()->pluck("unique_id");//おそらく結構無駄なことをしているため後程作り直す
   		$group = ($unique[0] % $groups[0]) + 1;
   		//以下は表示用のデータ取得

   		$contents = self::getContent($group, true);

   		return view("bbs", ["unique" => $unique, "id" => $id, "group" => $group , "contents" => $contents["contents"], "contents_id" => $contents["contents_id"], "hash" => $hash_val]);
		//return view("bbs", ["unique" => $unique, "id" => $id, "group" => $group , "contents" => $contents]);
   }
   public function postContent($id, $group, $content){
   	//DB::table("bbs")->create([('id' => $id, 'group_id' => $group, 'contents' => $contents)]);
   DB::table('bbs')->insert(['id' => $id, 'group_id' => $group, 'contents' => $content]);
	//DB::insert("Insert into table-name (id,group_id, contents) values (?,?,?)", [$id, $group, $content]);
   echo "<script type='text/javascript'>window.close();</script>";
   }

}
