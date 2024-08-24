"use client";

import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";
import { useRouter } from "next/navigation";
import ScheduleList from "../components/ScheduleList";

export default function MyPage() {
  const router = useRouter();

  // トークンを取得
  const token = Cookies.get("token");
  const userID = getUserIdFromToken(token);

  const handleCreateNew = () => {
    router.push("/create"); // 新規作成ページのパスを指定
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <h1>マイページ：{userID}</h1>
      <ScheduleList userId={userID} />
      <button onClick={handleCreateNew}>新規作成</button>
    </main>
  );
}

function getUserIdFromToken(token) {
  try {
    const decodedToken = jwtDecode(token);
    return decodedToken.user_id;
  } catch (error) {
    console.error("トークンのデコードに失敗しました:", error);
    return null;
  }
}
