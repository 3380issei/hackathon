import ScheduleCard from './ScheduleCard';

const ScheduleList = ({ schedules }) => {
    return (
        <div className="space-y-4">
            {schedules.map(schedule => (
                <ScheduleCard key={schedule.id} schedule={schedule} />
            ))}
        </div>
    );
};

export default ScheduleList;
