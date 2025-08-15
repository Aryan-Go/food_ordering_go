import React, { useState, useEffect } from "react";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import { useNavigate, useSearchParams } from "react-router-dom";

const Payment = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const [data, set_data] = useState([]);
  const [tip, set_tip] = useState();
  const id = searchParams.get("payment_id");
  useEffect(() => {
    const get_data = async () => {
      const data_sent = {
        id:Number(id)
      };
      const response = await axios.post(
        "/admin/admin_payment_complete",
        data_sent
      );
      set_data(response.data);
    };
    get_data();
  }, [tip]);
  if (data.status_code == 400) {
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
    //   navigate(-1)
  } else {
    return (
      <>
        <div className="h-[100vh] w-[100vw] bg-[url(payment_back.webp)] bg-no-repeat bg-black/50 bg-blend-overlay bg-size-[length:100%_100%] text-center flex flex-col items-center gap-2rem">
          <div className="mb-3">
            <h1 className="text-[5rem] font-bold text-green-500">
              Payment details
            </h1>
            <label
              htmlFor="exampleInputEmail1"
              className="htmlForm-label text-4xl font-bold"
            >
              For order_id
            </label>
            <input
              type="number"
              value={data.Order_id}
              readOnly
              min="0"
              className="htmlForm-control text-[2rem] px-[1rem] w-[10rem] text-green-500 font-bold outline-none"
              id="exampleInputEmail1"
              aria-describedby="emailHelp"
              name="order_id"
            />
          </div>
          <div className="mb-3">
            <label
              htmlFor="exampleInputEmail1"
              className="htmlForm-label htmlForm-label text-4xl font-bold"
            >
              Total Payment
            </label>
            <input
              type="number"
              value={data.Final_payment}
              readOnly
              min="0"
              className="htmlForm-control text-[2rem] px-[1rem] w-[10rem] text-green-500 font-bold outline-none"
              id="exampleInputEmail1"
              aria-describedby="emailHelp"
              name="total"
            />
          </div>
        </div>
      </>
    );
  }
};

export default Payment;
