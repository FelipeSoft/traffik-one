"use client"

import { BrainCircuit, MonitorDot, Route, ServerCog } from "lucide-react";
import Link from "next/link";
import { useParams, usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";

type Props = {
    children: React.ReactNode;
}

export const MainLayout = ({ children }: Props) => {
    const route = usePathname()

    return (
        <main className="flex flex-col w-full h-screen">
            <header className="p-6 flex items-center gap-2 z-30 border-b">
                <div className="rounded-full h-6 w-6 bg-gradient-to-b from-blue-300 to-blue-600"></div>
                <h1 className="font-semibold">Traffik<span className="text-blue-600">One</span></h1>
            </header>
            <div className="flex w-full h-full relative">
                <nav className="w-[250px] h-full bg-white/10 backdrop-blur-md z-20 border-r">
                    <ul>
                        <li className="my-4 rounded-md w-full flex">
                            <Link className={`${route === "/" && "text-blue-500"} w-full h-full py-2 px-6 flex items-center gap-4 hover:text-blue-500 transition-all`} href="/">
                                <MonitorDot />
                                Monitoring
                            </Link>
                        </li>
                        <li className="my-4 rounded-md w-full flex">
                            <Link className={`${route === "/backends" && "text-blue-500"} w-full h-full py-2 px-6 flex items-center gap-4 hover:text-blue-500 transition-all`} href="/backends">
                                <ServerCog />
                                Backends
                            </Link>
                        </li>
                        <li className="my-4 rounded-md w-full flex">
                            <Link className={`${route === "/routing-rules" && "text-blue-500"} w-full h-full py-2 px-6 flex items-center gap-4 hover:text-blue-500 transition-all`} href="/routing-rules">
                                <Route />
                                Routing Rules
                            </Link>
                        </li>
                        <li className="my-4 rounded-md w-full flex">
                            <Link className={`${route === "/algorithms" && "text-blue-500"} w-full h-full py-2 px-6 flex items-center gap-4 hover:text-blue-500 transition-all`} href="/algorithms">
                                <BrainCircuit />
                                Algorithms
                            </Link>
                        </li>
                    </ul>
                </nav>
                <section className="p-6 w-full h-full">
                    {children}
                </section>
            </div>
        </main>
    )
}   