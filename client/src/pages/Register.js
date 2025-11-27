import React, { useState } from "react";
import { api } from "../api/axios";
import { useNavigate, Link } from "react-router-dom";

function Register() {
  const navigate = useNavigate();

  const [step, setStep] = useState(1);
  const [message, setMessage] = useState("Hi! What is your name?");
  const [input, setInput] = useState("");

  const [user, setUser] = useState({
    name: "",
    email: "",
    password: "",
  });

  const handleNext = async (e) => {
    e.preventDefault();

    if (step === 1) {
      setUser({ ...user, name: input });
      setMessage("Nice! Now enter your email");
      setInput("");
      setStep(2);
    } 
    else if (step === 2) {
      setUser({ ...user, email: input });
      setMessage("Great! Set your password 🔐");
      setInput("");
      setStep(3);
    } 
    else if (step === 3) {
      const updatedUser = { ...user, password: input };
      setUser();
      setMessage("Registering your account…");
      
      try {
        await api.post("/register", updatedUser);
        setMessage("Account created successfully!");
        setTimeout(() => navigate("/"), 1500);
      } catch (err) {
        if (err.response?.data?.error) {
          setMessage(err.response.data.error);
        } else {
          setMessage("Something went wrong. Try again.");
        }
      }
    }
  };

  return (
    <div className="flex justify-center items-center bg-[#1c1c1c] h-screen text-center text-white">
      <div className="bg-[#101012] rounded-2xl p-10 w-[350px]">
        <div className="text-left mb-6">
          <p className="text-lg">{message}</p>
        </div>

        {step <= 3 && (
          <form onSubmit={handleNext} className="flex flex-col gap-y-5">
            <input className="p-3 border rounded-xl text-white bg-[#0e0e0f]" type={step === 3 ? "password" : "text"} placeholder={step === 1 ? "Your Name" : step === 2 ? "Your Email" : "Your Password"}
                  value={input} onChange={(e) => setInput(e.target.value)} required/>

            <button className="bg-black text-white font-bold px-4 py-2 rounded-xl mx-auto" type="submit">Next →</button>
          </form>
        )}
        <div className="mt-6">
          <p>Already have an account?</p>
          <Link to="/" className="text-white font-bold underline">Login Here</Link>
        </div>
      </div>
    </div>
  );
}

export default Register;

// import React, { useState } from 'react'
// import { api } from '../api/axios'
// import { useNavigate,Link } from 'react-router-dom'

// function Register() {

//   const [email,setEmail]=useState("")
//   const [password,setPassword]=useState("")
//   const [name,setName]=useState("")
//   const navigate = useNavigate()

//   const handleRegister = async (e)=>{
//     e.preventDefault();
//     try {
//       api.post("/register",{name,email,password})
//       navigate("/")
//     } catch (err) {
//             if (err.response && err.response.data?.error) {
//                 alert(err.response.data.error);
//             } else {
//                 alert("Something went wrong");
//             }
//     }
//   }
//   return (
//     <div className='flex justify-center items-center bg-[#1c1c1c] h-screen text-center'>
//        <div className=' bg-[#101012] rounded-2xl h-max w-max p-10'>
//         <form onSubmit={handleRegister}  className='flex flex-col gap-y-5 text-center w-max'>
//             <input className='p-3 border rounded-xl text-white bg-[#0e0e0f]' type='text' value={name} placeholder='Enter the Name' onChange={(e)=>setName(e.target.value)}/>
//             <input className='p-3 border rounded-xl text-white bg-[#0e0e0f]' type='email' value={email} placeholder='Set Your Email' onChange={(e)=>setEmail(e.target.value)}/>
//             <input className='p-3 border rounded-xl text-white bg-[#0e0e0f]' type='password' value={password} placeholder='Set Your Password' onChange={(e)=>setPassword(e.target.value)}/>
//             <button className='bg-[#000000] text-white font-bold w-max px-4 py-2 rounded-xl mx-auto' type='submit'>Register</button>
//         </form>
//         <div className='mt-5 text-white flex flex-col gap-y-2'>
//             <p>Already Have an account</p>
//             <Link to="/"className='text-white font-bold underline'>LogIn</Link>
//         </div>
//        </div>
//     </div>
//   )
// }

// export default Register
