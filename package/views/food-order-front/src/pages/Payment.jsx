import React, { useState, useEffect } from "react";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import { useNavigate, useSearchParams } from "react-router-dom";

const Payment = () => {
    const navigate = useNavigate()
    const [searchParams, setSearchParams] = useSearchParams();
    const [data, set_data] = useState([])
    const [tip, set_tip] = useState()
    const id = searchParams.get("order_id");
    const handlePayment = async() => {
        const data_sent = {
                order_id: Number(id),
                tip : Number(tip)
            };
            const response = await axios.post(
              "/customer/complete_payment",
              data_sent
            );
        toast.success("Payment has been done", {
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
        navigate("/customer");
    }
    useEffect(() => {
        const get_data = async () => {
            const data_sent = {
                order_id: Number(id),
                tip : Number(tip)
            };
            const response = await axios.post("/customer/render_payment", data_sent)
            set_data(response.data)
        }
        get_data()
    }, [tip])
    if (data.status_code == 400) {
        navigate(-1)
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
    }
    else {
        
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
                  value={id}
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
              <div className="dropdown">
                <label
                  htmlFor="tip"
                  className="htmlForm-label text-4xl font-bold"
                >
                  Tip
                </label>
                <input
                  value={tip}
                  onChange={(e) => {
                    set_tip(e.target.value);
                  }}
                  type="number"
                  min="0"
                  className="htmlForm-control text-[2rem] px-[1rem] w-[10rem] text-green-500 font-bold outline-none"
                  name="tip"
                />
              </div>
              <button onClick={handlePayment} type="button" className="btn btn-primary bg-green-500 text-white p-[2rem] w-[20rem] m-[1rem] rounded-2xl font-bold text-2xl">
                Done payment
              </button>
            </div>
          </>
        );
    }
}

export default Payment