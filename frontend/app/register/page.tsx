"use client"

import React, { useState } from "react";

export default function regisPage() {
    const [formData, setFormData] = useState({
        "username": "",
        "email": "",
        "password": "",
    });

    const handleRegis = async (e: React.FormEvent) => {
        e.preventDefault();

        const res = await fetch("http://localhost:8080/api/register", {
            method: "POST",
            headers: { "content-Type": "application/json" },
            body: JSON.stringify(formData),
        });

        const result = await res.json();
        if (res.ok) {
            alert("register berhasil");
            window.location.href = "/login";
        } else {
            alert("gagal: " + result.message);
        }
    };

    return (
        <div className="flex flex-col min-h-screen bg-white">
            <header className="fixed top-0 left-0 w-full h-[92px] md:h-[115px] bg-white flex items-center pl-[20px] md:pl-[127px] border-b-[4px] border-[#7E7F97]">
                <nav className="flex flex-col items-center pt-[12px] pb-[13px] md:pt-0 pb-0">
                        <a href=""><img src="/logo.png" alt="logo" className="w-[67px] h-[67px] md:w-[115px] md:h-[115px] object-contain"/></a>
                </nav>
            </header>

            <main className="flex flex-col flex-grow items-center mt-[132px] md:mt-[200px] ">
                <div className="w-[259px] md:w-[368px]">
                    <h1 className="font-bold text-[28px] md:text-[32px]">Daftar</h1>
                    <p className="text-[12px] pt-[10px]">Mari bergabung dengan kami!</p>
                </div>

                <form action="" className="flex flex-col w-[259px] md:w-[368px] pt-[40px]"> 
                    <label className="text-[15px] pb-[10px]">Nama</label>
                    <input type="text" name="nama" placeholder="Masukan nama" className="border-1 w-full pt-[12px] pb-[11px] rounded-[15px] pl-[15px] placeholder:text-[20px] placeholder:text-[#655F5F]"/>
                    <label className="text-[15px] pb-[10px] pt-[20px]">Email</label>
                    <input type="email" name="email" id="" placeholder="Masukan email" className="border-1 w-full pt-[12px] pb-[11px] pl-[15px] rounded-[15px] placeholder:text-[20px] placeholder:text-[#655F5F]"/>
                    <label className="text-[15px] pb-[10px] pt-[20px]">Password</label>
                    <input type="password" name="password" id="" placeholder="Masukan password" className="border-1 w-full pt-[12px] pb-[11px] pl-[15px] rounded-[15px] placeholder:text-[20px] placeholder:text-[#655F5F]"/>
                    <button type="submit" className="bg-[#125E9C] pt-[14px] pb-[13px] rounded-[15px] mt-[55px] text-white border-1 border-black cursor-pointer hover:bg-[#133C5D]">Daftar</button>
                    <div className="flex justify-center pt-[60px]">
                        <p>Sudah memiliki akun? <a href="" className="text-[#6781D9] hover:text-[#213578]">Masuk</a></p>
                    </div>
                </form>
            </main>

            <footer className="w-full h-[400px] border-t-[4px] border-[#7E7F97] mt-[70px] flex flex-col md:flex-row">
                <div className="pl-[40px] pt-[30px] md:pl-[128px] md:pt-[44px]">
                    <img src="/logo.png" alt="logo" className="w-[66px] h-[66px] md:w-[115] md:h-[115px]"/>
                    <p className="text-[#5C5858] text-[19px] pt-[20px] md:pt-0">Belajar bersama di kursku</p>
                </div>

                <div className="pl-[40px] pt-[50px] md:pt-[84px] md:pl-[103px]">
                    <h3 className="font-bold text-[21px]">Tentang kami</h3>
                    <a href=""><p className="text-[19px] pt-[40px]">tentang kursku</p></a>
                </div>
            </footer>
        </div>
    );
}