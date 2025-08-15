import React,{useEffect} from 'react'
import Navbar from "../components/Navbar_customer.jsx";
import axios from "../axios/AxiosHandler.jsx";

const Customer_chef = () => {
  useEffect(() => {
    const send_data = async () => {
      await axios.post("/customer/cus_chef");
    }
    send_data()
  })
  return (
    <>
      <Navbar />
      <div class=" text-center flex flex-col items-center justify-center h-[100vh] bg-[url(customer_home2.webp)] bg-no-repeat bg-size-[length:100vw_100vh]">
        <h1 class="text-4xl text-green-500 text-center font-bold m-[20rem] text-shadow-black">
          Your request for conversion to a chef has been sent to our admin.
          Please enjoy your time as a customer till the admin accepts your
          request
        </h1>
      </div>
    </>
  );
}

export default Customer_chef