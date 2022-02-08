import React, { useEffect, useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Nav from './components/Nav';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import cookie from 'react-cookies';

function App() {
  const [name,setName]=useState('');
  useEffect(()=>{
    (
        async ()=>{
          const response= await fetch('http://localhost:8000/data/user',{
                headers:{'Content-Type':'application/json','token':cookie.load("token")},
                credentials:'include',
          })

          const content=await response.json();
          if (content!=null){
              if (content.err_code===0){
                  setName(content.data.name);
              }else{
                  console.log(content.msg)
              }
          }

        }

    )();
});

  return (
    <div className="App">
      
        <BrowserRouter>
        <Nav name={name} setName={setName} />

      <main className="form-signin">
        <Routes>
          <Route path="/" element={<Home name={name} />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login setName={setName} />} />
        </Routes>
      </main>
      
      </BrowserRouter>
        
    </div>
  );
}

export default App;
