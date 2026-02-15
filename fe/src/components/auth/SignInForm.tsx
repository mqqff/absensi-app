import { useState } from "react";
import { EyeCloseIcon, EyeIcon } from "../../icons";
import Label from "../form/Label";
import Input from "../form/input/InputField";
import Button from "../ui/button/Button";
import Swal from 'sweetalert2'
import {useNavigate} from "react-router";

export default function SignInForm() {
    const navigate = useNavigate();

    const [showPassword, setShowPassword] = useState(false);
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [loading, setLoading] = useState(false);

    const BASE_URL = import.meta.env.VITE_API_BASE_URL;

    const handleSubmit = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        setLoading(true);

        try {
            const res = await fetch(`${BASE_URL}/auth/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "x-api-key": import.meta.env.VITE_API_KEY,
                },
                body: JSON.stringify({
                    email,
                    password,
                }),
            });

            const data = await res.json();

            if (!res.ok) {
                await Swal.fire({
                    icon: "error",
                    title: "Oops...",
                    text: "Email atau password salah!",
                });
            }

            if (data.payload.access_token != null || data.payload.access_token != "") {
                localStorage.setItem("token", data.payload.access_token);

                const response = await fetch(`${BASE_URL}/auth/session`, {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                        "x-api-key": import.meta.env.VITE_API_KEY,
                    }
                });

                const session = await response.json();

                localStorage.setItem("user", JSON.stringify(session.payload));

                navigate("/");
            }
        } catch (err) {

        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="flex flex-col flex-1">
            <div className="flex flex-col justify-center flex-1 w-full max-w-md mx-auto">
                <div>
                    <div className="mb-5 sm:mb-7">
                        <h1 className="mb-2 font-semibold text-gray-800 text-title-sm dark:text-white/90 sm:text-title-md">
                            Sign In
                        </h1>
                    </div>

                    <form onSubmit={handleSubmit}>
                        <div className="space-y-6">
                            <div>
                                <Label>Email <span className="text-error-500">*</span></Label>
                                <Input
                                    placeholder="info@gmail.com"
                                    value={email}
                                    required
                                    onChange={(e) => setEmail(e.target.value)}
                                />
                            </div>

                            <div>
                                <Label>Password <span className="text-error-500">*</span></Label>
                                <div className="relative">
                                    <Input
                                        type={showPassword ? "text" : "password"}
                                        placeholder="Masukkan password"
                                        required
                                        value={password}
                                        onChange={(e) => setPassword(e.target.value)}
                                    />
                                    <span
                                        onClick={() => setShowPassword(!showPassword)}
                                        className="absolute z-30 -translate-y-1/2 cursor-pointer right-4 top-1/2"
                                    >
                    {showPassword ? (
                        <EyeIcon className="fill-gray-500 dark:fill-gray-400 size-5" />
                    ) : (
                        <EyeCloseIcon className="fill-gray-500 dark:fill-gray-400 size-5" />
                    )}
                  </span>
                                </div>
                            </div>

                            <div>
                                <Button className="w-full" size="sm" disabled={loading}>
                                    {loading ? "Signing in..." : "Sign in"}
                                </Button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}
