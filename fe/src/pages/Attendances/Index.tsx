import { useEffect, useState } from "react";
import PageBreadcrumb from "../../components/common/PageBreadCrumb";
import { useNavigate } from "react-router";
import PageMeta from "../../components/common/PageMeta";
import ComponentCard from "../../components/common/ComponentCard";
import Label from "../../components/form/Label";
import Swal from "sweetalert2";

export default function Index() {
    const token = localStorage.getItem("token");
    const navigate = useNavigate();

    const [now, setNow] = useState(new Date());
    const [loading, setLoading] = useState(false);

    const [attendance, setAttendance] = useState<{
        check_in?: string;
        check_out?: string;
        status?: number;
    } | null>(null);

    useEffect(() => {
        const timer = setInterval(() => {
            setNow(new Date());
        }, 1000);
        return () => clearInterval(timer);
    }, []);

    const fetchTodayAttendance = async () => {
        try {
            const res = await fetch(
                `${import.meta.env.VITE_API_BASE_URL}/attendances/open`,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                        "X-API-KEY": import.meta.env.VITE_API_KEY,
                    },
                }
            );

            const data = await res.json();

            setAttendance(data.payload ?? null);
        } catch (err) {
            console.error(err);
        }
    };

    useEffect(() => {
        fetchTodayAttendance();
    }, []);

    const handleCheckIn = async () => {
        try {
            setLoading(true);

            const res = await fetch(
                `${import.meta.env.VITE_API_BASE_URL}/attendances/checkin`,
                {
                    method: "POST",
                    headers: {
                        Authorization: `Bearer ${token}`,
                        "X-API-KEY": import.meta.env.VITE_API_KEY,
                    },
                }
            );

            const data = await res.json();

            if (!res.ok) {
                await Swal.fire("Gagal", data.payload.error.message || "Gagal check in", "error");
                return;
            }

            await Swal.fire("Berhasil", "Check In berhasil", "success");
            fetchTodayAttendance();
        } catch (err: any) {
            Swal.fire("Error", err.message, "error");
        } finally {
            setLoading(false);
        }
    };

    const handleCheckOut = async () => {
        try {
            setLoading(true);

            const res = await fetch(
                `${import.meta.env.VITE_API_BASE_URL}/attendances/checkout`,
                {
                    method: "PATCH",
                    headers: {
                        Authorization: `Bearer ${token}`,
                        "X-API-KEY": import.meta.env.VITE_API_KEY,
                    },
                }
            );

            const data = await res.json();

            if (!res.ok) {
                await Swal.fire("Gagal", data.payload.error.message || "Gagal check out", "error");
                return;
            }

            await Swal.fire("Berhasil", "Check Out berhasil", "success");
            fetchTodayAttendance();
        } catch (err: any) {
            await Swal.fire("Error", err.message, "error");
        } finally {
            setLoading(false);
        }
    };

    const formatTime = (dateString?: string) => {
        if (!dateString) return "-";
        return new Date(dateString).toLocaleTimeString("id-ID");
    };

    const canCheckIn = !attendance?.check_in;
    const canCheckOut = attendance?.check_in && !attendance?.check_out;

    return (
        <div>
            <PageMeta
                title="Presensi Hari Ini"
                description="Halaman presensi pegawai"
            />
            <PageBreadcrumb pageTitle="Presensi" />

            <div className="grid grid-cols-1 gap-6 max-w-xl">
                <ComponentCard title="Presensi Hari Ini">
                    <div className="space-y-4">

                        <div>
                            <Label>Waktu Sekarang</Label>
                            <div className="text-lg font-semibold">
                                {now.toLocaleTimeString("id-ID")}
                            </div>
                        </div>

                        <div>
                            <Label>Check In</Label>
                            <div className="text-base">
                                {formatTime(attendance?.check_in)}
                            </div>
                        </div>

                        <div>
                            <Label>Check Out</Label>
                            <div className="text-base">
                                {formatTime(attendance?.check_out)}
                            </div>
                        </div>

                        <div className="pt-4">
                            {canCheckIn && (
                                <button
                                    onClick={handleCheckIn}
                                    disabled={loading}
                                    className="w-full px-4 py-2 text-white bg-green-600 rounded-lg hover:bg-green-700 disabled:opacity-50"
                                >
                                    {loading ? "Memproses..." : "Check In"}
                                </button>
                            )}

                            {canCheckOut && (
                                <button
                                    onClick={handleCheckOut}
                                    disabled={loading}
                                    className="w-full px-4 py-2 text-white bg-blue-600 rounded-lg hover:bg-blue-700 disabled:opacity-50"
                                >
                                    {loading ? "Memproses..." : "Check Out"}
                                </button>
                            )}

                            {!canCheckIn && !canCheckOut && (
                                <div className="text-center text-green-600 font-medium">
                                    Presensi hari ini selesai ✔
                                </div>
                            )}
                        </div>

                    </div>
                </ComponentCard>
            </div>
        </div>
    );
}
