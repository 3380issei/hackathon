import { jwtDecode } from "jwt-decode";

export const fetchUserByID = async (id) => {
  try {
    const response = await fetch(`http://localhost:8080/users/${id}`);
    if (!response.ok) {
      throw new Error("ユーザー情報の取得に失敗しました");
    }
    const data = await response.json();
    console.log("ユーザー情報を取得しました:", data);
    return data;
  } catch (error) {
    console.error("エラーが発生しました:", error);
  }
};

export const getUserIDFromToken = (token) => {
  try {
    const decodedToken = jwtDecode(token);
    return decodedToken.user_id;
  } catch (error) {
    console.error("トークンのデコードに失敗しました:", error);
    return null;
  }
};
