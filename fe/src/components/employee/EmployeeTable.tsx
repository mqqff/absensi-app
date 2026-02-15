import {
    Table,
    TableBody,
    TableCell,
    TableHeader,
    TableRow,
} from "../ui/table";
import { useEffect, useState } from "react";
import Badge from "../ui/badge/Badge";
import {
    Employee,
    positionLabel,
    departmentLabel,
    statusLabel,
} from "../../types/employee.ts";
import {useNavigate} from "react-router";
import Swal from "sweetalert2";

export default function EmployeeTable() {
    const navigate = useNavigate();

    const [employees, setEmployees] = useState<Employee[]>([]);
    const [page, setPage] = useState(1);
    const [totalPage, setTotalPage] = useState(1);
    const [limit] = useState(10);

    const [filters, setFilters] = useState({
        identifier: "",
        position: "",
        department: "",
        status: 0,
    });

    const token = localStorage.getItem("token");

    useEffect(() => {
        const params = new URLSearchParams({
            page: String(page),
            limit: String(limit),
            ...(filters.identifier && { identifier: filters.identifier }),
            ...(filters.position && { position: filters.position }),
            ...(filters.department && { department: filters.department }),
            ...(filters.status > 0 && { status: String(filters.status) }),
        });

        fetch(`${import.meta.env.VITE_API_BASE_URL}/employees?${params}`, {
            headers: {
                Authorization: `Bearer ${token}`,
                "X-API-KEY": import.meta.env.VITE_API_KEY,
            },
        })
            .then((res) => res.json())
            .then((res) => {
                setEmployees(res.payload?.employees ?? []);
                setTotalPage(res.payload?.meta?.total_page ?? 1);
                console.log(res.payload.employees)
            })
            .catch(console.error);
    }, [page, filters]);

    const handleDelete = async (id: string) => {
        const confirm = await Swal.fire({
            title: "Hapus pegawai?",
            text: "Data tidak bisa dikembalikan",
            icon: "warning",
            showCancelButton: true,
            confirmButtonText: "Ya, hapus",
            cancelButtonText: "Batal",
        });

        if (!confirm.isConfirmed) return;

        try {
            const res = await fetch(
                `${import.meta.env.VITE_API_BASE_URL}/employees/${id}`,
                {
                    method: "DELETE",
                    headers: {
                        Authorization: `Bearer ${token}`,
                        "X-API-KEY": import.meta.env.VITE_API_KEY,
                    },
                }
            );

            const data = await res.json();
            if (!res.ok) throw new Error(data.message);

            setEmployees((prev) => prev.filter((e) => e.id !== id));

            await Swal.fire("Berhasil", "Pegawai dihapus", "success");
        } catch (err: any) {
            await Swal.fire("Error", err.message, "error");
        }
    };


    return (
        <div className="overflow-hidden rounded-xl border border-gray-200 bg-white dark:border-white/[0.05] dark:bg-white/[0.03]">
            <div className="flex justify-end p-4 border-b border-gray-100 dark:border-white/[0.05]">
                <button
                    onClick={() => navigate("/employees/create")}
                    className="px-4 py-2 text-sm text-white bg-blue-600 rounded-lg hover:bg-blue-700"
                >
                    + Tambah Pegawai
                </button>
            </div>
            {/* FILTER */}
            <div className="grid grid-cols-1 md:grid-cols-5 gap-3 p-4 border-b border-gray-100 dark:border-white/[0.05]">
                <input
                    placeholder="Cari nama / email..."
                    autoFocus
                    className="px-3 py-2 border rounded-lg text-sm"
                    value={filters.identifier}
                    onChange={(e) => {
                        setPage(1);
                        setFilters({ ...filters, identifier: e.target.value });
                    }}
                />

                <select
                    className="px-3 py-2 border rounded-lg text-sm"
                    value={filters.position}
                    onChange={(e) => {
                        setPage(1);
                        setFilters({ ...filters, position: e.target.value });
                    }}
                >
                    <option value="">Semua Posisi</option>
                    {Object.entries(positionLabel).map(([key, label]) => (
                        <option key={key} value={key}>
                            {label}
                        </option>
                    ))}
                </select>

                <select
                    className="px-3 py-2 border rounded-lg text-sm"
                    value={filters.department}
                    onChange={(e) => {
                        setPage(1);
                        setFilters({ ...filters, department: e.target.value });
                    }}
                >
                    <option value="">Semua Departemen</option>
                    {Object.entries(departmentLabel).map(([key, label]) => (
                        <option key={key} value={key}>
                            {label}
                        </option>
                    ))}
                </select>

                <select
                    className="px-3 py-2 border rounded-lg text-sm"
                    value={filters.status || ""}
                    onChange={(e) => {
                        setPage(1);
                        setFilters({
                            ...filters,
                            status: e.target.value ? Number(e.target.value) : 0,
                        });
                    }}
                >
                    <option value="">Semua Status</option>
                    {Object.entries(statusLabel).map(([key, label]) => (
                        <option key={key} value={Number(key)}>
                            {label}
                        </option>
                    ))}
                </select>

                <button
                    onClick={() => {
                        setPage(1);
                        setFilters({
                            identifier: "",
                            position: "",
                            department: "",
                            status: 0,
                        });
                    }}
                    className="px-3 py-2 border rounded-lg text-sm"
                >
                    Reset
                </button>
            </div>

            {/* TABLE */}
            <div className="max-w-full overflow-x-auto">
                <Table>
                    <TableHeader className="border-b border-gray-100 dark:border-white/[0.05]">
                        <TableRow>
                            <TableCell isHeader className="px-5 py-3 font-medium text-gray-500 text-start text-theme-xs dark:text-gray-400">
                                Nama Pegawai
                            </TableCell>

                            <TableCell isHeader className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400">
                                Departemen
                            </TableCell>

                            <TableCell isHeader className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400">
                                Kontak
                            </TableCell>

                            <TableCell isHeader className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400">
                                Status
                            </TableCell>

                            <TableCell isHeader className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400">
                                Gaji
                            </TableCell>

                            <TableCell isHeader className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400">
                                Aksi
                            </TableCell>
                        </TableRow>
                    </TableHeader>

                    <TableBody className="divide-y divide-gray-100 dark:divide-white/[0.05]">
                        {employees.length === 0 && (
                            <TableRow>
                                <TableCell colSpan={5} className="text-center py-6 text-gray-500">
                                    Data tidak ditemukan
                                </TableCell>
                            </TableRow>
                        )}

                        {employees.map((emp) => (
                            <TableRow key={emp.id}>
                                <TableCell className="px-5 py-4 sm:px-6 text-start">
                                    <div>
                    <span className="block font-medium text-gray-800 text-theme-sm dark:text-white/90">
                      {emp.name}
                    </span>
                                        <span className="block text-gray-500 text-theme-xs dark:text-gray-400">
                      {emp.position}
                    </span>
                                    </div>
                                </TableCell>

                                <TableCell className="px-4 py-3 text-gray-500 text-center text-theme-sm dark:text-gray-400">
                                    {emp.department}
                                </TableCell>

                                <TableCell className="px-4 py-3 text-gray-500 text-center text-theme-sm dark:text-gray-400">
                                    <div>
                                        <div>{emp.email}</div>
                                        <div className="text-theme-xs">{emp.phone}</div>
                                    </div>
                                </TableCell>

                                <TableCell className="px-4 py-3 text-center">
                                    <Badge
                                        size="sm"
                                        color={
                                            emp.status === "Aktif"
                                                ? "success"
                                                : emp.status === "Tidak Aktif"
                                                    ? "error"
                                                    : "warning"
                                        }
                                    >
                                        {emp.status}
                                    </Badge>
                                </TableCell>

                                <TableCell className="px-4 py-3 text-gray-500 text-theme-sm dark:text-gray-400 text-center">
                                    Rp {emp.salary.toLocaleString("id-ID")}
                                </TableCell>

                                <TableCell className="px-4 py-3 text-center">
                                    <div className="flex justify-center gap-2">
                                        <button
                                            onClick={() => navigate(`/employees/${emp.id}/edit`)}
                                            className="px-3 py-1 text-xs text-white bg-yellow-500 rounded-full font-bold hover:bg-yellow-600"
                                        >
                                            Ubah
                                        </button>

                                        <button
                                            onClick={() => handleDelete(emp.id)}
                                            className="px-3 py-1 text-xs text-white bg-red-600 rounded-full font-bold hover:bg-red-700"
                                        >
                                            Hapus
                                        </button>
                                    </div>
                                </TableCell>

                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>

            {/* PAGINATION */}
            <div className="flex items-center justify-between px-6 py-4 border-t border-gray-100 dark:border-white/[0.05]">
                <button
                    onClick={() => setPage((p) => Math.max(1, p - 1))}
                    disabled={page === 1}
                    className="px-3 py-1 text-sm border rounded-lg disabled:opacity-50"
                >
                    Previous
                </button>

                <span className="text-sm text-gray-500">
          Page {page} of {totalPage}
        </span>

                <button
                    onClick={() => setPage((p) => Math.min(totalPage, p + 1))}
                    disabled={page === totalPage}
                    className="px-3 py-1 text-sm border rounded-lg disabled:opacity-50"
                >
                    Next
                </button>
            </div>
        </div>
    );
}
