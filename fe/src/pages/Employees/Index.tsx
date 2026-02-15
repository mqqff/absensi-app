import PageBreadcrumb from "../../components/common/PageBreadCrumb";
import ComponentCard from "../../components/common/ComponentCard";
import PageMeta from "../../components/common/PageMeta";
import EmployeeTable from "../../components/employee/EmployeeTable";

export default function Index() {
    return (
        <>
            <PageMeta
                title="Kelola Pegawai"
                description="Halaman untuk mengelola pegawai"
            />
            <PageBreadcrumb pageTitle="Semua Pegawai" />
            <div className="space-y-6">
                <ComponentCard title="Tabel Pegawai">
                    <EmployeeTable />
                </ComponentCard>
            </div>
        </>
    );
}
