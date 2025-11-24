import React,{useState} from 'react'
import NavBar from '../components/NavBar';
import { api } from '../api/axios';
import { LuUpload } from "react-icons/lu";


function AdminHome() {

  const [text, setText] = useState("");
  const [loading,setLoading]=useState(false)
  const name = localStorage.getItem("name")

  const uploadText = async(e)=>{
      e.preventDefault();
      try {
        setText("");
        await api.post("/upload",{text});
        alert("Material Uploaded Successfully")
      } catch (error) {
        console.log(error);
        alert(error.response?.data?.error || "Something went wrong");
      }
  }

  return (
    <div className='min-h-screen bg-black text-white'>
      <NavBar/>
      <div>
        <h2 className='text-center text-3xl tracking-wider pt-16 lg:pt-0 p-5 lg:text-5xl'>HI🙋‍♂️ , {name}</h2>
        <div className='flex flex-col gap-y-10 lg:mt-10 mt-24'>
          <textarea className='p-5 text-white bg-[#1b1b1e]  w-[90%] mx-auto rounded-2xl lg:w-[50%] h-56 lg:h-96' value={text} onChange={(e) => setText(e.target.value)} placeholder="Paste study notes here"/>
          <button className='border flex items-center justify-center gap-x-3 lg:text-2xl w-max mx-auto px-4 py-2 font-bold rounded-3xl' onClick={uploadText}><LuUpload /><p>{loading?"Uploading...":"Upload"}</p></button>
        </div>
      </div>
    </div>
  )
}

export default AdminHome
