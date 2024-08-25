import Link from "next/link";

export default function Header() {
    return (
        <header className="w-full p-4 mb-12 bg-red-500 text-white text-center">
            <Link href="/">
                <h1 className="text-4xl font-bold cursor-pointer">起床外出圧力くん</h1>
            </Link>
        </header>
    );
}
