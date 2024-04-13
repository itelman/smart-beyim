//  import <HTMLDivElement> to the ref type
import React, { PureComponent, useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  PolarGrid,
  CartesianAxis,
} from "recharts";

import FeatherIcon from "feather-icons-react";

import { cn } from "../../utils/utils";

import FilterDate from "../UI/FilterDate";

import { ICustomLineChartData } from "../UI/assets";

interface CustomLineChartProps {
  data: ICustomLineChartData[];
}

interface ICustomLineChartProps
  extends CustomLineChartProps,
    React.HTMLAttributes<HTMLDivElement> {}

const CustomLineChart: React.FC<ICustomLineChartProps> = ({
  data,
  ...props
}) => {
  const [dateFrom, setDateFrom] = useState<string>("thisWeek");
  const [currentSkill, setCurrentSkill] = useState<string>("Reading");

  const options = ["Reading", "Listening", "Speaking", "Writing"];
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);
  const minAxisYInterval =
    data.reduce((min, p) => (p.score < min ? p.score : min), data[0].score) -
    0.5;

  const handleDateChange = (value: string) => {
    setDateFrom(value);
  };

  return (
    <div
      className={cn(
        "rounded-[12px] bg-white px-[20px] py-[21px]",
        props.className,
      )}
    >
      <header className="mb-[36px] flex items-center justify-between">
        <p className="text-subHeading3  font-semibold">Results</p>
        <div className="relative">
          <div
            className={`flex cursor-pointer items-center gap-[7px] text-black-100 `}
            onClick={() => setIsDropdownOpen(!isDropdownOpen)}
            onBlur={() => setIsDropdownOpen(false)}
          >
            <p className="text-label1">{currentSkill}</p>
            <FeatherIcon
              icon="chevron-down"
              className="text-label1"
              size={16}
            />
          </div>
          {/*  dropdown  */}
          {isDropdownOpen && (
            <div className="absolute right-0 top-[120%] z-10 flex flex-col rounded-md   bg-white  py-3 shadow-2xl">
              {options.map((option) => (
                <div
                  key={option}
                  onClick={() => {
                    handleDateChange(option);
                    setIsDropdownOpen(false);
                  }}
                  className={cn(
                    `flex w-[130px] cursor-pointer items-center gap-[7px] px-5 py-[6px] text-black-300   hover:bg-primary-700 hover:text-white ${option === currentSkill && "bg-primary-700 text-white"}`,
                  )}
                >
                  <p className="text-label1">{option}</p>
                </div>
              ))}
            </div>
          )}
        </div>
      </header>

      <ResponsiveContainer width="100%" height={447}>
        <LineChart
          width={500}
          height={300}
          data={data}
          margin={{
            top: 5,
            right: 30,
            // left: 20,
            bottom: 30,
          }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          {/* <CartesianAxis axisLine={true} /> */}
          <XAxis dataKey="date" axisLine={false} dy={20} tickLine={false} />
          <YAxis
            // interval={"preserveStart"}
            domain={[minAxisYInterval, 9]}
            axisLine={false}
            tickLine={false}
            dx={-25}
          />
          <Tooltip />
          <Line dataKey="score" stroke="#4379EE" />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};

export default CustomLineChart;
