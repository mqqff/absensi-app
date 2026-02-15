export interface Employee {
    id: string;
    name: string;
    email: string;
    phone: string;
    position: string;
    department: string;
    salary: number;
    address: string;
    status: string;
    created_at: string;
    updated_at: string;
    deleted_at: string | null;
}

export const positionLabel: Record<number, string> = {
    2: "HR",
    3: "Admin",
    4: "Employee"
};

export const departmentLabel: Record<number, string> = {
    1: "HRD",
    2: "Engineering / IT",
    3: "Keuangan",
    4: "Operasional",
    5: "Sales & Marketing",
};

export const statusLabel: Record<number, string> = {
    2: "Aktif",
    1: "Tidak Aktif",
};

export const optionPosition = [
    { value: "2", label: "HR" },
    { value: "3", label: "Admin" },
    { value: "4", label: "Employee" },
];

export const optionDepartment = [
    { value: "1", label: "HRD" },
    { value: "2", label: "Engineering / IT" },
    { value: "3", label: "Keuangan" },
    { value: "4", label: "Operasional" },
    { value: "5", label: "Sales & Marketing" },
];

export const optionStatus = [
    { value: "1", label: "Aktif" },
    { value: "2", label: "Tidak Aktif" },
];