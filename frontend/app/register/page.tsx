"use client"

import React, { useState } from "react";
import { useMutation } from "@tanstack/react-query";
// 
export default function RegisPage() {
const [formData, setFormData] = useState({
    nama: "",
    email: "",
    password: ""
})
// 
const mutation = useMutation({
    mutationFn: async (newAccount) => {
        const response = await fetch("http://localhost:8080/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newAccount),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || "kesalahan pada server");
        }

        return response.json();
    },
    onSuccess: (data) => {
        console.log("berhasil login", data)
    },
    onError: (error) => {
        alert(error.message);
    }
});
// 
const isFormEmpty = !formData.nama || !formData.email || !formData.password;
// 
const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
        ...formData,
        [e.target.name]: e.target.value
    });
}
const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault
}
    return (
        <div className="flex flex-col min-h-screen bg-white">
            <header className="fixed top-0 left-0 w-full h-[92px] md:h-[92px] bg-white flex items-center pl-[20px] md:pl-[70px] border-b-[4px] border-[#7E7F97]">
                <nav className="flex flex-col items-center pt-[12px] pb-[13px] md:pt-0 pb-0">
                        <a href=""><img src="/logo.png" alt="logo" className="w-[67px] h-[67px] md:w-[92px] md:h-[92px] object-contain"/></a>
                </nav>
            </header>

            <main className="flex flex-col flex-grow items-center mt-[132px] md:mt-[162px] ">
                <div className="w-[259px] md:w-[368px]">
                    <h1 className="font-bold text-[28px] md:text-[32px]">Daftar</h1>
                    <p className="text-[12px] pt-[10px]">Mari bergabung dengan kami!</p>
                </div>

                <form onSubmit={handlesub} action="" className="flex flex-col w-[259px] md:w-[368px] pt-[40px]"> 
                    <label className="text-[15px] pb-[10px]">Nama</label>
                    <input type="text" name="nama" placeholder="Masukan nama" className="border-1 w-full pt-[12px] pb-[11px] rounded-[15px] pl-[15px] placeholder:text-[20px] placeholder:text-[#655F5F] md:placeholder:text-[18px]"/>
                    <label className="text-[15px] pb-[10px] pt-[20px]">Email</label>
                    <input type="email" name="email" id="" placeholder="Masukan email" className="border-1 w-full pt-[12px] pb-[11px] pl-[15px] rounded-[15px] placeholder:text-[20px] placeholder:text-[#655F5F] md:placeholder:text-[18px]"/>
                    <label className="text-[15px] pb-[10px] pt-[20px]">Password</label>
                    <input type="password" name="password" id="" placeholder="Masukan password" className="border-1 w-full pt-[12px] pb-[11px] pl-[15px] rounded-[15px] placeholder:text-[20px] placeholder:text-[#655F5F] md:placeholder:text-[18px]"/>
                    <button type="submit" className="bg-[#125E9C] pt-[14px] pb-[13px] rounded-[15px] mt-[55px] text-white border-1 border-black cursor-pointer hover:bg-[#133C5D] md:text-[20px]">Daftar</button>
                    <div className="flex justify-center pt-[60px] md:text-[20px]">
                        <p>Sudah memiliki akun? <a href="" className="text-[#6781D9] hover:text-[#213578]">Masuk</a></p>
                    </div>
                </form>
            </main>

            <footer className="w-full h-[400px] border-t-[4px] border-[#7E7F97] mt-[70px] flex flex-col md:flex-row">
                <div className="pl-[40px] pt-[30px] md:pl-[70px] md:pt-[20px]">
                    <img src="/logo.png" alt="logo" className="w-[66px] h-[66px] md:w-[78px] md:h-[78px]"/>
                    <p className="text-[#5C5858] text-[19px] md:text-[17px] pt-[20px] md:pt-[20px]">Belajar bersama di kursku</p>
                </div>

                <div className="pl-[40px] pt-[50px] md:pt-[39px] md:pl-[103px]">
                    <h3 className="font-bold text-[21px] md:text-[20px]">Tentang kami</h3>
                    <a href=""><p className="text-[19px] md:text-[18px] pt-[40px] md:pt-[30px]">tentang kursku</p></a>
                </div>
            </footer>
        </div>
    );
}