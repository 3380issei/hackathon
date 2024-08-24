"use client";

import { useState } from "react";
import LoginForm from "./components/LoginForm";
import SignupForm from "./components/SignupForm";

export default function Home() {
  const [isLogin, setIsLogin] = useState(true);

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <div className="w-full max-w-sm">
        {isLogin ? <LoginForm /> : <SignupForm />}
        <br></br>
        <div className="flex justify-center mb-6">
          <button
            className={`px-4 py-2 mr-2 ${
              isLogin ? "bg-blue-500 text-white" : "bg-gray-200"
            } rounded`}
            onClick={() => setIsLogin(true)}
          >
            ログイン
          </button>
          <button
            className={`px-4 py-2 ${
              !isLogin ? "bg-blue-500 text-white" : "bg-gray-200"
            } rounded`}
            onClick={() => setIsLogin(false)}
          >
            新規登録
          </button>
        </div>
      </div>
    </main>
  );
}
