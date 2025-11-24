import React, { useState } from 'react'
import { api } from '../api/axios'
import { useNavigate,Link } from 'react-router-dom'

function Register() {

  const [email,setEmail]=useState("")
  const [password,setPassword]=useState("")
  const [name,setName]=useState("")
  const navigate = useNavigate()

  const handleRegister = async (e)=>{
    e.preventDefault();
    try {
      api.post("/register",{name,email,password})
      navigate("/")
    } catch (err) {
            if (err.response && err.response.data?.error) {
                alert(err.response.data.error);
            } else {
                alert("Something went wrong");
            }
    }
  }

  return (
    <div className='flex justify-center items-center bg-[#D78FEE] h-screen text-center'>
       <div className=' bg-[#9B5DE0] rounded-2xl h-max w-max p-10'>
        <form onSubmit={handleRegister}  className='flex flex-col gap-y-5 text-center w-max'>
            <input className='p-3 rounded-xl' type='text' value={name} placeholder='Enter the Name' onChange={(e)=>setName(e.target.value)}/>
            <input className='p-3 rounded-xl' type='text' value={email} placeholder='Set Your Email' onChange={(e)=>setEmail(e.target.value)}/>
            <input className='p-3 rounded-xl' type='password' value={password} placeholder='Set Your Password' onChange={(e)=>setPassword(e.target.value)}/>
            <button className='bg-purple-950 text-white font-bold w-max px-4 py-2 rounded-xl mx-auto' type='submit'>Register</button>
        </form>
        <div className='mt-5 text-white flex flex-col gap-y-2'>
            <p>Already Have an account</p>
            <Link to="/"className='text-purple-950 font-bold underline'>LogIn</Link>
        </div>
       </div>
    </div>
  )
}

export default Register
