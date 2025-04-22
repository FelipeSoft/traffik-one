"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { Button } from "@/components/ui/button"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"

const FormSchema = z.object({
    username: z.string({
        required_error: "Required"
    }),
    password: z.string({
        required_error: "Required"
    })
})

export default function Auth() {
    const form = useForm<z.infer<typeof FormSchema>>({
        resolver: zodResolver(FormSchema),
        defaultValues: {
            username: "",
            password: ""
        },
    })

    function onSubmit(data: z.infer<typeof FormSchema>) {

    }

    return (
        <main className="flex flex-col w-full h-screen justify-center items-center">
            <header className="p-4 flex items-center gap-2 z-30">
                <div className="rounded-full h-6 w-6 bg-gradient-to-b from-blue-300 to-blue-600"></div>
                <h1 className="font-semibold">Traffik<span className="text-blue-600">One</span></h1>
            </header>
            <div className="flex justify-center items-center relative">
                <div className="absolute inset-0 -z-10">
                    <div className="rounded-full h-20 w-20 bg-gradient-to-b from-blue-300 to-blue-600 -bottom-12 -left-[40px] absolute"></div>
                    <div className="rounded-full h-32 w-32 bg-gradient-to-b from-blue-300 to-blue-600 bottom-12 left-[280px] absolute"></div>
                </div>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="w-[350px] space-y-6 p-6 bg-white/10 backdrop-blur-md z-20 shadow-md border rounded-md">
                        <h2 className="text-center text-neutral-700 text-xl">Welcome to TraffikOne</h2>
                        <FormField
                            control={form.control}
                            name="username"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Username</FormLabel>
                                    <FormControl>
                                        <Input placeholder="john.doe" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="password"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Password</FormLabel>
                                    <FormControl>
                                        <Input type="password" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <Button type="submit" className="bg-blue-500 w-full cursor-pointer">Submit</Button>
                    </form>
                </Form>
            </div>
        </main>

    );
}
