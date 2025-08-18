import React, { useState, useEffect } from "react";
import axios from "../axios/AxiosHandler.jsx";
import Navbar from "../components/Navbar_customer.jsx";
import { toast, Bounce } from "react-toastify";
import { useNavigate } from "react-router-dom";
const Menu = () => {
    const navigate = useNavigate()
    let id_arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    const [quant_arr,set_quanr_arr] = useState([0,0,0,0,0,0,0,0,0])
    const [special_instruc,set_special_instruc] = useState(["","","","","","","","",""])
  function enforceMinMax(el) {
    if (el.value != "") {
      if (parseInt(el.value) < parseInt(el.min)) {
        el.value = el.min;
      } else if (!isNumber(el.value)) {
        el.value = el.min;
      }
    }
  }
  const [data, setData] = useState([]);
  const [data2, setData2] = useState([]);
//   const [quant, setQuant] = useState();
//   const [spec, setSpec] = useState();
  useEffect(() => {
    const getMenu = async () => {
      let response = await axios.get("/customer/menu_show");
      // setData(response.data.slice(0, 9));
      setData2(response.data)
    };
    getMenu();
  }, []);
    console.log(data2);
    const handleSpex = (e, id) => {
        const updated = [...special_instruc];
        updated[id - 1] = e.target.value;
        set_special_instruc(updated);
    }
    const handleQuant = (e, id) => {
        const updated = [...quant_arr];
        updated[id-1] = Number(e.target.value);
        set_quanr_arr(updated);
    };
    const handleClick = async() => {
        const data = {
          item_add: quant_arr,
            instructions: special_instruc,
          id:id_arr
        };
        const response = await axios.post("/customer/food_items_added", data);
        if (response.data.status_code == 401) {
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
          toast.success("Your order is = " + response.data.message, {
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
          navigate(`/waiting_page?order_id=${response.data.message}`);
        }
    }
  if (data.status_code == 401) {
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
    useEffect(() => {
      const getMenu = async () => {
        let response = await axios.get("/customer/menu_show");
        setData(response.data.slice(0, 9));
        // setData2(response.data);
        response = null;
      };
      getMenu();
    }, []);
    const menu_table = data.map((value, key) => {
      return (
        <tr key={key} className="m-[2rem] p-[2rem]">
          <td className="text-xl font-bold mx-[2rem]">{value.name}</td>
          <td className="text-xl font-semibold mx-[2rem] ">
            {value.description}
          </td>
          <td className="text-xl font-bold mx-[2rem]">Rs{value.price}</td>
          <td className="text-xl font-bold mx-[2rem]">
            <label htmlFor="quant">Quanity</label>
            <input
              pattern="^[1-9][0-9]*$"
              required
              type="number"
              value={quant_arr[value.id - 1]}
              onChange={(e) => handleQuant(e, value.id)}
              name="quant"
              min="0"
              onKeyUp={(e) => enforceMinMax(e.target)}
              className="quant user-invalid:border-red-500 w-[5rem] text-center border-dashed border-black"
            />
            <input
              readOnly
              type="number"
              value={value.id}
              name="food_id"
              className="food_id"
              hidden
            />
          </td>
          <td className="text-xl font-bold mx-[2rem]">
            <label htmlFor="quant">Special Instructions here</label>
            <br />
            <input
              type="text"
              name="special_instructions"
              className="food_id"
              value={special_instruc[value.id - 1]}
              onChange={(e) => handleSpex(e, value.id)}
            />
          </td>
        </tr>
      );
    });
    return (
      <>
        <Navbar />
        <form
          id="myForm"
          action="/food_items_added"
          method="post"
                className="text-center"
                onSubmit={(e) => {e.preventDefault()}}
        >
          <div className="h-full w-full bg-[url(menu_back.webp)] bg-no-repeat bg-size-[length:100%_100%] text-center flex flex-col justify-center items-center gap-2rem text-center">
            <h1 className="text-[6rem] font-bold">Menu</h1>
            <table className="text-center w-[70vw] m-[4rem] px-[10rem]">
              <thead>
                <tr className="m-[1rem]">
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Item
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Description
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Price
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Quantity Added
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Special Instructions
                  </th>
                </tr>
              </thead>
              <tbody>{menu_table}</tbody>
            </table>
            <button
              type="submit"
              className="text-center text-4xl font-bold w-[30%] bg-[#F4D871] px-[2rem] py-[0.5rem] mb-30"
              onClick={handleClick}
            >
              Submit your order
            </button>
          </div>
        </form>
      </>
    );
  }
};

export default Menu;
