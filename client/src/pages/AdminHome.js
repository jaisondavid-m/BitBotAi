import React,{useState} from 'react'

function AdminHome() {

  const [text, setText] = useState("");

    const uploadText = async () => {
    await fetch("http://localhost:8000/upload", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ text }),
    });
    alert("Material uploaded");
  };

  return (
    <div>
      <h1>AI Study Assistant</h1>

      <hr />

      <h2>Teacher Upload</h2>
      <textarea rows={6} cols={60} value={text} onChange={(e) => setText(e.target.value)} placeholder="Paste study notes here"/>
      <br />
      <button onClick={uploadText}>Upload</button>
      <hr />
    </div>
  )
}

export default AdminHome
