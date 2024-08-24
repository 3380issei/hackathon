// src/components/ScheduleList.jsx

import React, { useEffect, useState } from 'react';
import ScheduleCard from './ScheduleCard';

const ScheduleList = ({ userId }) => {
    const [schedules, setSchedules] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchSchedules = async () => {
            try {
                const response = await fetch(`http://localhost:8080/schedules/${userId}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch schedules');
                }
                const data = await response.json();
                setSchedules(data);
            } catch (error) {
                setError(error.message);
            } finally {
                setLoading(false);
            }
        };

        fetchSchedules();
    }, [userId]);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="space-y-4">
            {schedules.map(schedule => (
                <ScheduleCard key={schedule.ID} schedule={schedule} />
            ))}
        </div>
    );
};

export default ScheduleList;
