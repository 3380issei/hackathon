export const fetchSchedulesByUserID = async (userID) => {
  try {
    const response = await fetch(`http://localhost:8080/schedules/${userID}`);
    if (!response.ok) {
      throw new Error("Failed to fetch schedules");
    }
    const data = await response.json();
    console.log("スケジュールを取得しました:", data);
    return data;
  } catch (error) {
    console.error("Error fetching schedules)", error);
  }
};
