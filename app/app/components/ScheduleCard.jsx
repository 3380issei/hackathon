import React from 'react';

const ScheduleCard = ({ schedule }) => {
    console.log(schedule);

    return (
        <div className="p-4 border border-gray-300 rounded shadow-md">
            <h2 className="text-xl font-bold">{schedule.destination}</h2>
            <p><strong>Expired:</strong> {schedule.expired ? 'Yes' : 'No'}</p>
            <p><strong>日時:</strong> {schedule.deadline}</p>
        </div>
    );
};

export default ScheduleCard;
