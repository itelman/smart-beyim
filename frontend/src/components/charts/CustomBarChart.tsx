//  import <HTMLDivElement> to the ref type
import React, { PureComponent, useState } from "react";

import FeatherIcon from "feather-icons-react";
import { ICustomBarChartData } from "../UI/assets/index";
import { cn } from "../../utils/utils";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";
import FilterDate from "../UI/FilterDate";

interface CustomBarChartProps {
  data: ICustomBarChartData[];
}

interface ICustomBarChartProps
  extends CustomBarChartProps,
    React.HTMLAttributes<HTMLDivElement> {}

const CustomBarChart: React.FC<ICustomBarChartProps> = ({ data, ...props }) => {
  const [dateFrom, setDateFrom] = useState<string>("thisWeek");

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
        <p className="text-subHeading3  font-semibold">Average Scores</p>
        <FilterDate dateFrom={dateFrom} handleDateChange={handleDateChange} />
      </header>

      <ResponsiveContainer width="100%" height={447}>
        <BarChart
          data={data}
          barSize={14}
          margin={{
            top: 5,
            // left: 20,
            bottom: 30,
          }}
        >
          <XAxis
            axisLine={false}
            tickLine={false}
            dataKey="short"
            padding={{ left: 23, right: 23 }}
            dy={19}
            fontSize={20}
          />
          <YAxis
            axisLine={false}
            tickLine={false}
            fontSize={20}
            dx={-10}
            width={30}
          />
          <Tooltip />
          <Bar
            dataKey="score"
            radius={50}
            fill="#55BBEB"
            background={{ fill: "#EEF0FA" }}
          />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default CustomBarChart;
