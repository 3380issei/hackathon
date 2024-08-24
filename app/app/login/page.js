"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        const result = await response.json();
        console.log("Login successful:", result);
        // ログイン成功後にダッシュボードなどのページにリダイレクト
        router.push("/dashboard"); // 例: ダッシュボードページ
      } else {
        const error = await response.json();
        console.error("Login error:", error);
        // エラーメッセージの表示などの処理を追加
      }
    } catch (error) {
      console.error("Network error:", error);
    }
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <form
        className="flex flex-col space-y-4 w-full max-w-sm"
        onSubmit={handleSubmit}
      >
        <label htmlFor="email" className="flex flex-col">
          Email
          <input
            type="email"
            id="email"
            name="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="p-2 border border-gray-300 rounded"
          />
        </label>
        <label htmlFor="password" className="flex flex-col">
          Password
          <input
            type="password"
            id="password"
            name="password"
            required
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="p-2 border border-gray-300 rounded"
          />
        </label>
        <button type="submit" className="p-2 bg-blue-500 text-white rounded">
          Login
        </button>
      </form>
    </main>
  );
}
