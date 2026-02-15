import {BrowserRouter as Router, Routes, Route} from "react-router";
import SignIn from "./pages/AuthPages/SignIn";
import NotFound from "./pages/OtherPage/NotFound";
import AppLayout from "./layout/AppLayout";
import {ScrollToTop} from "./components/common/ScrollToTop";
import Home from "./pages/Dashboard/Home";
import EmployeeIndex from "./pages/Employees/Index"
import EmployeeCreate from "./pages/Employees/Create.tsx";
import EmployeeEdit from "./pages/Employees/Edit.tsx";
import AttendanceIndex from "./pages/Attendances/Index.tsx";

export default function App() {
    return (
        <>
            <Router>
                <ScrollToTop/>
                <Routes>
                    {/* Dashboard Layout */}
                    <Route element={<AppLayout/>}>
                        <Route index path="/" element={<Home/>}/>

                        <Route path="/employees" element={<EmployeeIndex/>}/>
                        <Route path="/employees/create" element={<EmployeeCreate/>}/>
                        <Route path="/employees/:id/edit" element={<EmployeeEdit/>} />

                        <Route path="/attendance" element={<AttendanceIndex />} />
                    </Route>

                    {/* Auth Layout */}
                    <Route path="/signin" element={<SignIn/>}/>

                    {/* Fallback Route */}
                    <Route path="*" element={<NotFound/>}/>
                </Routes>
            </Router>
        </>
    );
}
