import { useState } from "react";

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
import {useNavigate} from "react-router";
import Swal from "sweetalert2";

export default function Create() {
    const navigate = useNavigate();
    const token = localStorage.getItem("token");

    const [loading, setLoading] = useState(false);

    const [form, setForm] = useState({
        name: "",
        email: "",
        phone: "",
        salary: "",
        password: "",
        position: "",
        department: "",
        status: "",
        address: "",
    });

    const handleChange = (field: string, value: string) => {
        setForm((prev) => ({ ...prev, [field]: value }));
    };

    const handleSubmit = async () => {
        try {
            setLoading(true);

            const res = await fetch(
                `${import.meta.env.VITE_API_BASE_URL}/employees`,
                {
                    method: "POST",
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
                        password: form.password,
                        position: Number(form.position),
                        department: Number(form.department),
                        status: Number(form.status),
                        address: form.address,
                    }),
                }
            );

            const data = await res.json();

            console.log(data.payload)

            if (!res.ok) {
                await Swal.fire({
                    icon: "error",
                    title: "Oops...",
                    text: data.payload.error.message || "pastikan seluruh kolom terisi dengan benar",
                });
                return
            }

            await Swal.fire({
                icon: "success",
                title: "Yeayy...",
                text: "berhasil menambah pegawai",
            });
            navigate("/employees");
        } catch (err: any) {
            alert(err.message);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <PageMeta
                title="Tambah Pegawai Baru"
                description="Halaman untuk menambahkan pegawai baru"
            />
            <PageBreadcrumb pageTitle="Tambah Pegawai" />

            <div className="grid grid-cols-1 gap-6 max-w-4xl">
                <ComponentCard title="Pegawai Baru">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">

                        <div>
                            <Label>Nama</Label>
                            <Input
                                type="text"
                                placeholder="Nama Pegawai"
                                value={form.name}
                                onChange={(e) => handleChange("name", e.target.value)}
                                autoFocus
                            />
                        </div>

                        <div>
                            <Label>Email</Label>
                            <Input
                                type="email"
                                placeholder="Email Pegawai"
                                value={form.email}
                                onChange={(e) => handleChange("email", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Phone</Label>
                            <Input
                                type="text"
                                placeholder="No. Telp Pegawai"
                                value={form.phone}
                                onChange={(e) => handleChange("phone", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Gaji</Label>
                            <Input
                                type="number"
                                placeholder="Gaji Pegawai"
                                value={form.salary}
                                onChange={(e) => handleChange("salary", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Password</Label>
                            <Input
                                type="password"
                                placeholder="Password"
                                value={form.password}
                                onChange={(e) => handleChange("password", e.target.value)}
                            />
                        </div>

                        <div>
                            <Label>Jabatan</Label>
                            <Select
                                options={optionPosition}
                                placeholder="Pilih Jabatan"
                                onChange={(value) => handleChange("position", value)}
                                className="dark:bg-dark-900"
                            />
                        </div>

                        <div>
                            <Label>Departemen</Label>
                            <Select
                                options={optionDepartment}
                                placeholder="Pilih Departemen"
                                onChange={(value) => handleChange("department", value)}
                                className="dark:bg-dark-900"
                            />
                        </div>

                        <div>
                            <Label>Status</Label>
                            <Select
                                options={optionStatus}
                                placeholder="Pilih Status"
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
                                type="submit"
                                disabled={loading}
                                className="px-5 py-2 text-white bg-blue-600 rounded-lg hover:bg-blue-700 disabled:opacity-50"
                            >
                                {loading ? "Menyimpan..." : "Simpan Pegawai"}
                            </button>
                        </div>
                    </div>
                </ComponentCard>
            </div>
        </div>
    );
}
