import React,{useState} from 'react'
import { api } from '../api/axios';

function AdminHome() {

  const [text, setText] = useState("");

  const uploadText = async(e)=>{
      e.preventDefault();
      try {
        await api.post("/upload",{text});
        alert("Material Uploaded Successfully")
      } catch (error) {
        console.log(error);
        alert(error.response?.data?.error || "Something went wrong");
      }
  }

  return (
    <div>
      <h1>AI Study Assistant</h1>
      <hr/>
      <h2>Teacher Upload</h2>
      <textarea rows={6} cols={60} value={text} onChange={(e) => setText(e.target.value)} placeholder="Paste study notes here"/>
      <br />
      <button onClick={uploadText}>Upload</button>
      <hr />
    </div>
  )
}

export default AdminHome
