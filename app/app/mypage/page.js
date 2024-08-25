"use client";

import { useState, useEffect } from "react";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";
import ScheduleList from "../components/ScheduleList";
import { fetchUserByID, getUserIDFromToken } from "../services/userService";
import { fetchSchedulesByUserID } from "../services/scheduleService";
import Header from "../components/Header";

export default function MyPage() {
  const router = useRouter();
  const [userID, setUserID] = useState(null);
  const [user, setUser] = useState(null);
  const [schedules, setSchedules] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const token = Cookies.get("token");
      if (token) {
        const id = getUserIDFromToken(token);
        setUserID(id);

        if (id) {
          try {
            // ユーザー情報の取得
            const fetchedUser = await fetchUserByID(id);
            setUser(fetchedUser);

            // スケジュールの取得
            const fetchedSchedules = await fetchSchedulesByUserID(id);
            setSchedules(fetchedSchedules);
          } catch (error) {
            console.error("データの取得に失敗しました", error);
          }
        }
      }
    };

    fetchData();
  }, []);

  const handleCreateNew = () => {
    router.push("/create"); // 新規作成ページのパスを指定
  };

  return (
    <>
      <Header />
      <main className="flex min-h-screen flex-col items-center justify-center p-24">
        {user ? (
          <div className="mb-6">
            <h1 className="text-2xl font-bold mb-6">
              {user.name}さんのマイページ
            </h1>
          </div>
        ) : (
          <p>ユーザー情報を取得中...</p>
        )}
        {schedules.length > 0 ? (
          <ScheduleList schedules={schedules} />
        ) : (
          <p>スケジュールはありません</p>
        )}
        <button
          onClick={handleCreateNew}
          className="px-6 py-3 mt-4 bg-red-600 text-white rounded-lg shadow-lg transition-transform transform hover:scale-105"
        >
          新規作成
        </button>
      </main>
    </>
  );
}
