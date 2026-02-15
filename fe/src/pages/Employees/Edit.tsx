import { useEffect, useState } from "react";

import PageBreadcrumb from "../../components/common/PageBreadCrumb";
import PageMeta from "../../components/common/PageMeta";
import Label from "../../components/form/Label.tsx";
import Input from "../../components/form/input/InputField.tsx";
import ComponentCard from "../../components/common/ComponentCard.tsx";
import Select from "../../components/form/Select.tsx";
import TextArea from "../../components/form/input/TextArea.tsx";

import {
    optionDepartment,
    optionPosition,
    optionStatus,
} from "../../types/employee.ts";

import { useNavigate, useParams } from "react-router";
import Swal from "sweetalert2";

export default function Edit() {
    const navigate = useNavigate();
    const { id } = useParams();
    const token = localStorage.getItem("token");

    const [loading, setLoading] = useState(false);
    const [loadingData, setLoadingData] = useState(true);

    const [form, setForm] = useState({
        name: "",
        email: "",
        phone: "",
        salary: "",
        position: "",
        department: "",
        status: "",
        address: "",
    });

    const handleChange = (field: string, value: string) => {
        setForm((prev) => ({ ...prev, [field]: value }));
    };

    // ✅ GET DETAIL EMPLOYEE
    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await fetch(
                    `${import.meta.env.VITE_API_BASE_URL}/employees/${id}`,
                    {
                        headers: {
                            Authorization: `Bearer ${token}`,
                            "X-API-KEY": import.meta.env.VITE_API_KEY,
                        },
                    }
                );

                const data = await res.json();
                const emp = data.payload;

                setForm({
                    name: emp.name ?? "",
                    email: emp.email ?? "",
                    phone: emp.phone ?? "",
                    salary: String(emp.salary ?? ""),
                    position: String(emp.position_idx ?? emp.position ?? ""),
                    department: String(emp.department_idx ?? emp.department ?? ""),
                    status: String(emp.status_idx ?? emp.status ?? ""),
                    address: emp.address ?? "",
                });
            } catch (err: any) {
                Swal.fire("Error", err.message, "error");
            } finally {
                setLoadingData(false);
            }
        };

        fetchData();
    }, [id]);

    // ✅ UPDATE EMPLOYEE
    const handleSubmit = async () => {
        try {
            setLoading(true);

            const res = await fetch(
                `${import.meta.env.VITE_API_BASE_URL}/employees/${id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${token}`,
                        "X-API-KEY": import.meta.env.VITE_API_KEY,
                    },
                    body: JSON.stringify({
                        name: form.name,
                        email: form.email,
                        phone: form.phone,
                        salary: Number(form.salary),
                        position: Number(form.position),
                        department: Number(form.department),
                        status: Number(form.status),
                        address: form.address,
                    }),
                }
            );

            const data = await res.json();

            if (!res.ok) {
                await Swal.fire({
                    icon: "error",
                    title: "Oops...",
                    text:
                        data?.payload?.error?.message ||
                        "Gagal memperbarui data",
                });
                return;
            }

            await Swal.fire({
                icon: "success",
                title: "Berhasil",
                text: "Data pegawai diperbarui",
            });

            navigate("/employees");
        } catch (err: any) {
            Swal.fire("Error", err.message, "error");
        } finally {
            setLoading(false);
        }
    };

    if (loadingData) {
        return <div className="p-6">Loading data...</div>;
    }

    return (
        <div>
            <PageMeta
                title="Edit Pegawai"
                description="Halaman untuk mengubah data pegawai"
            />
            <PageBreadcrumb pageTitle="Edit Pegawai" />

            <div className="grid grid-cols-1 gap-6 max-w-4xl">
                <ComponentCard title="Edit Pegawai">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">

                        <div>
                            <Label>Nama</Label>
                            <Input
                                type="text"
                                value={form.name}
                                onChange={(e) => handleChange("name", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Email</Label>
                            <Input
                                type="email"
                                value={form.email}
                                onChange={(e) => handleChange("email", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Phone</Label>
                            <Input
                                type="text"
                                value={form.phone}
                                onChange={(e) => handleChange("phone", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Gaji</Label>
                            <Input
                                type="number"
                                value={form.salary}
                                onChange={(e) => handleChange("salary", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Jabatan</Label>
                            <Select
                                options={optionPosition}
                                defaultValue={form.position}
                                onChange={(value) => handleChange("position", value)}
                                className="dark:bg-dark-900"
                            />
                        </div>

                        <div>
                            <Label>Departemen</Label>
                            <Select
                                options={optionDepartment}
                                defaultValue={form.department}
                                onChange={(value) => handleChange("department", value)}
                                className="dark:bg-dark-900"
                            />
                        </div>

                        <div>
                            <Label>Status</Label>
                            <Select
                                options={optionStatus}
                                defaultValue={form.status}
                                onChange={(value) => handleChange("status", value)}
                                className="dark:bg-dark-900"
                            />
                        </div>

                        <div className="md:col-span-2">
                            <Label>Alamat</Label>
                            <TextArea
                                value={form.address}
                                onChange={(value) => handleChange("address", value)}
                                rows={2}
                            />
                        </div>

                        <div className="md:col-span-2 flex justify-end pt-2">
                            <button
                                onClick={handleSubmit}
                                disabled={loading}
                                className="px-5 py-2 text-white bg-blue-600 rounded-lg hover:bg-blue-700 disabled:opacity-50"
                            >
                                {loading ? "Menyimpan..." : "Update Pegawai"}
                            </button>
                        </div>
                    </div>
                </ComponentCard>
            </div>
        </div>
    );
}
