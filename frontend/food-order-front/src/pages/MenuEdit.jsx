import React, { useState } from "react";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar_admin.jsx";

const MenuEdit = () => {
  const [name, SetName] = useState("");
  const [desc, SetDesc] = useState("");
  const [price, SetPrice] = useState("");
  const [category, SetCategory] = useState("");
  const role = "customer";
  const navigate = useNavigate();
  const clickHandler = async (name, desc, price, category) => {
    const data = {
      id: 0,
      name: name,
      description: desc,
      price: Number(price),
      c_id: Number(category),
    };
    const response = await axios.post("/admin/menu_edit", data);
    if (response.data.status_code == 400) {
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
  return (
    <>
      <Navbar />
      <div className="w-[100vw] h-[100vh] bg-size-[100%_100%] grayscale-30 bg-[url(https://imgs.search.brave.com/TouscCUt008bVNGiF-OauhIv-4eqphkTGUWPWSRV25g/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/ZnJlZXBpay5jb20v/cHJlbWl1bS12ZWN0/b3Ivd29ybGQtZm9v/ZC1kYXktYmFja2dy/b3VuZF83MjYyMzct/Mjk4LmpwZz9zZW10/PWFpc19oeWJyaWQm/dz03NDA)] bg-no-repeat bg-center flex justify-center items-center">
        <div className="w-max p-[2rem] bg-[#372A36] rounded-2xl">
          <h1 className="text-6xl font-bold bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
            Add menu items
          </h1>
          <form
            className="flex flex-col justify-center items-center"
            onSubmit={(e) => e.preventDefault()}
          >
            <label
              htmlFor="name"
              className="text-white font-bold text-[1.5rem]"
            >
              Food Name
            </label>
            <input
              name="name"
              type="text"
              id="name"
              className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]"
              value={name}
              onChange={(e) => SetName(e.target.value)}
            />
            <label
              htmlFor="text"
              className="text-white font-bold text-[1.5rem]"
            >
              Description
            </label>
            <textarea
              name="text"
              className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]"
              value={desc}
              onChange={(e) => SetDesc(e.target.value)}
            />
            <label
              htmlFor="price"
              className="text-white font-bold text-[1.5rem]"
            >
              Price
            </label>
            <input
              name="price"
              type="text"
              className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]"
              value={price}
              onChange={(e) => SetPrice(e.target.value)}
            />
            <label
              htmlFor="category"
              className="text-white font-bold text-[1.5rem]"
            >
              Category
            </label>
            <select
              value={category}
              onChange={(e) => SetCategory(e.target.value)}
              name="order_id"
              className="text-[1.5rem] w-max text-white border-2 border-white rounded-2xl p-[0.5rem] mt-[1rem]"
            >
              <option className="text-white" value={1}>
                Starters
              </option>
              <option className="text-white" value={2}>
                Main Course
              </option>
              <option className="text-white" value={3}>
                Desert
              </option>
            </select>
            <div>
              <button
                type="submit"
                className="bg-gray-500 px-[1rem] rounded-2xl text-white m-[1rem] text-bold text-[2rem] my-[2rem]"
                onClick={() => clickHandler(name, desc, price, category)}
              >
                Submit
              </button>
            </div>
          </form>
        </div>
      </div>
    </>
  );
};

export default MenuEdit;
