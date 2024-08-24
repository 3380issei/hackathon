import { useState } from 'react';

export default function SignupForm() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();

        try {
            const response = await fetch('http://localhost:8080/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password, name }),
            });

            if (response.ok) {
                const result = await response.json();
                console.log('Success:', result);
                // 登録成功後にログイン画面にリダイレクト
                window.location.href = '/login'; // または window.location.replace('/login');
            } else {
                const error = await response.json();
                console.error('Error:', error);
                // エラーメッセージの表示などの処理を追加
            }
        } catch (error) {
            console.error('Network error:', error);
            // ネットワークエラーの処理を追加
        }
    };

    return (
        <form className="flex flex-col space-y-4 w-full max-w-sm" onSubmit={handleSubmit}>
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
            <label htmlFor="name" className="flex flex-col">
                名前
                <input
                    type="text"
                    id="name"
                    name="name"
                    required
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    className="p-2 border border-gray-300 rounded"
                />
            </label>
            <button type="submit" className="p-2 bg-blue-500 text-white rounded">
                Submit
            </button>
        </form>
    );
}
