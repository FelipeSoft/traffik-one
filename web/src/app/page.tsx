import { MainLayout } from "@/components/layouts/main-layout";
import Link from "next/link";

export default function Home() {
    return (
        <MainLayout>
            <ul className="w-full h-full grid grid-cols-4">
                <li className="border rounded-md h-[300px] w-[300px]">
                    <Link href="/">OK!</Link>
                </li>
            </ul>
        </MainLayout>
    );
}
