import React,{useEffect,useState} from 'react'
import Navbar from "../components/Navbar_customer.jsx";
import axios from "../axios/AxiosHandler.jsx";
import { useNavigate } from "react-router-dom";
import { toast, Bounce } from "react-toastify";

const Customer_chef = () => {
  const [data,set_data] = useState([])
  useEffect(() => {
    const send_data = async () => {
      const resp = await axios.post("/customer/cus_chef");
      set_data(resp.data)
    }
    send_data()
  },[])
  if (data.status_code == 202) {
    return (
      <>
        <Navbar />
        <div class="min-h-[100vh] text-center flex flex-col items-center justify-center h-[100vh] bg-[url(customer_home2.webp)] bg-no-repeat bg-size-[length:100vw_100vh]">
          <h1 class="text-4xl text-green-500 text-center font-bold m-[20rem] text-shadow-black">
            Your request for conversion to a chef has been sent to our admin.
            Please enjoy your time as a customer till the admin accepts your
            request
          </h1>
        </div>
      </>
    );
  }
  else {
     toast.error(data.message, {
       position: "top-center",
       autoClose: 5000,
       hideProgressBar: false,
       closeOnClick: false,
       pauseOnHover: true,
       draggable: true,
       progress: undefined,
       theme: "dark",
       transition: Bounce,
     });
    if (data.message == "Malformed Token") {
      setTimeout(() => {
        navigate("/login");
      }, 2000);
    } else if (
      data.message == "This is a protected route where only customer is allowed"
    ) {
      setTimeout(() => {
        navigate(-1);
      }, 2000);
    }
  }
}

export default Customer_chef