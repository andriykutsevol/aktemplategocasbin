import { createContext, useContext, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { useLocalStorage } from "./useLocalStorage";
const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useLocalStorage("user", null);
  const [token, setToken] = useLocalStorage("access_token", null);
  const navigate = useNavigate();

  // call this function when you want to authenticate the user
  const login = async (data) => {

    //--------------------------------------------------------

    //const url = 'http://localhost:8080/api/v1/pub/login';
    //const url = 'http://template_go_react_golang:8080/api/v1/pub/login';
    const url = '/api/v1/pub/login';

    const payload = {
        captcha_code: "captcha_code",
        captcha_id: "captcha_id",
        password: data.password,
        user_name: data.username,
      };

      try {
        const res = await fetch(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'accept': 'application/json',
          },
          body: JSON.stringify(payload),
        });
  
        if (!res.ok) {
          throw new Error('Network response was not ok');
        }
  
        const resdata = await res.json();

        //console.log(resdata)
        setUser(data.username);
        setToken(resdata.access_token);
        //console.log(window.localStorage.getItem("access_token"));

      } catch (error) {
        setError(error);
      }

    //--------------------------------------------------------      

    navigate("/weather");
  };

  // call this function to sign out logged in user
  const logout = () => {
    setUser(null);
    navigate("/login", { replace: true });
  };

  const value = useMemo(
    () => ({
      user,
      login,
      logout,
    }),
    [user]
  );
  
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
  return useContext(AuthContext);
};