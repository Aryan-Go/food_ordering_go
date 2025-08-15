import React, { useState, useEffect } from "react";
import { toast, Bounce } from "react-toastify";
import { useNavigate } from "react-router-dom";
import axios from "../axios/AxiosHandler.jsx";

const Admin = () => {
  const navigate = useNavigate();
    const [data, set_data] = useState(null);
    const [select_order_id , set_select_order_id] = useState()
    const [select_payment_id , set_select_payment_id] = useState()
    const [select_customer_id , set_select_customer_id] = useState()
  useEffect(() => {
    const get_data = async () => {
      const response = await axios.get("/admin/admin_details");
      set_data(response.data);
    };
    get_data();
  }, []);
useEffect(() => {
  const get_data_2 = async () => {
    if (data.incomplete_order?.length > 0) {
      set_select_order_id(data.incomplete_order[0]);
    }
    if (data.incomplete_payment?.length > 0) {
      set_select_payment_id(data.incomplete_payment[0]);
    }
    if (data.chef_customer?.length > 0) {
      set_select_customer_id(data.chef_customer[0]);
    }
  };
  get_data_2();
}, [data]);
    const IncompleteOrder = (order_id) => {
      console.log(order_id)
    navigate(`/waiting_page?order_id=${order_id}`);
};
const IncompletePayment = (payment_id) => {
        navigate(`/payment_admin?payment_id=${payment_id}`);

    }
  console.log(data);
  const CustomerToChef = async (id) => {
    const data_sent = {
      id: id,
    };
    const response = await axios.post(
      "/admin/admin_chef_conversion",
      data_sent
    );
    if (data.status_code == 401) {
      toast.error(response.data.message, {
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
    } else {
      toast.success(response.data.message, {
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
  };
  if (data == null) {
    return (
      <>
        <div className="h-[100vh]">
          <h1 className="py-[2rem]">Hello Admin</h1>
          <h2 className="py-[2rem]">Loading ... </h2>
        </div>
      </>
    );
  } else {
      if (data.status_code) {
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
           return (
           <>
             <div className="h-[100vh]">
               <h1 className="py-[1rem] text-[8rem] font-bold text-red-500">You are not allowed</h1>
             </div>
           </>
         );
      }
      else {
          return (
            <div className="h-max flex flex-col items-center w-[100vw] m-[1rem]">
              <h1 className="py-[1rem] text-[8rem] font-bold">Hello Admin</h1>
              <h2 className="py-[2rem] text-[4rem] font-bold">
                Order and its Preperation
              </h2>
              <form
                onSubmit={(e) => {
                  e.preventDefault();
                }}
              >
                <label
                  htmlFor="quant"
                  className="px-[4rem] font-bold text-[2rem]"
                >
                  Order_id
                </label>
                {data.incomplete_order == null ? (
                  <h1 className="text-red-500 text-4xl font-bold ">
                    No data for now
                  </h1>
                ) : (
                  <>
                    <select
                      value={select_order_id}
                      onChange={(e) => set_select_order_id(e.target.value)}
                      name="order_id"
                      className="text-[2rem] w-max"
                    >
                      {data.incomplete_order.map((value, key) => {
                        console.log(value);
                        return <option value={value}>{value}</option>;
                      })}
                    </select>
                    <br />
                    <br />
                    <button
                      onClick={() => IncompleteOrder(Number(select_order_id))}
                      type="button"
                      className="text-center text-4xl font-bold w-max bg-[#F4D871] px-[2rem] py-[0.5rem] mb-30"
                    >
                      Find this order
                    </button>
                  </>
                )}
              </form>

              <h2 className="py-[2rem] text-[4rem] font-bold">Payment_table</h2>
              <form
                onSubmit={(e) => {
                  e.preventDefault();
                }}
              >
                <label
                  htmlFor="quant"
                  className="px-[4rem] font-bold text-[2rem]"
                >
                  Order_id
                </label>
                {data.incomplete_payment == null ? (
                  <h1 className="text-red-500 text-4xl font-bold ">
                    No data for now
                  </h1>
                ) : (
                  <>
                    <select
                      value={select_payment_id}
                      onChange={(e) => set_select_payment_id(e.target.value)}
                      name="order_id"
                      className="text-[2rem] w-max"
                    >
                      {data.incomplete_payment.map((value, key) => {
                        return <option value={value}>{value}</option>;
                      })}
                    </select>
                    <br />
                    <br />
                    <button
                      onClick={() => IncompletePayment(Number(select_payment_id))}
                      type="button"
                      className="text-center text-4xl font-bold w-max bg-[#F4D871] px-[2rem] py-[0.5rem] mb-30"
                    >
                      Find this payment
                    </button>
                  </>
                )}
              </form>

              <h2 className="py-[2rem] text-[4rem] font-bold">
                Convert to chef
              </h2>
              <form
                onSubmit={(e) => {
                  e.preventDefault();
                }}
              >
                <label
                  htmlFor="quant"
                  className="px-[4rem] font-bold text-[2rem]"
                >
                  Order_id
                </label>
                {data.chef_customer == null ? (
                  <h1 className="text-red-500 text-4xl font-bold ">
                    No data for now
                  </h1>
                ) : (
                  <>
                    <select
                      value={select_customer_id}
                      onChange={(e) => set_select_customer_id(e.target.value)}
                      name="order_id"
                      className="text-[2rem] w-max"
                    >
                      {data.chef_customer.map((value, key) => {
                        return <option value={value}>{value}</option>;
                      })}
                    </select>
                    <br />
                    <br />
                    <button
                      onClick={() => CustomerToChef(Number(select_customer_id))}
                      type="button"
                      className="text-center text-4xl font-bold w-max bg-[#F4D871] px-[2rem] py-[0.5rem] mb-30"
                    >
                      Convert this customer
                    </button>
                  </>
                )}
              </form>
            </div>
          );
      }
    
  }
};

export default Admin;
