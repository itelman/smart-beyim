import { useEffect, useState } from "react";
import { CustomPieChart } from "../../components/charts/CustomPieChart";
import CustomBarChart from "../../components/charts/CustomBarChart";
import RecentOrders from "../../components/UI/RecentOrders";
import CustomLineChart from "../../components/charts/CustomLineChart";
import { ITransaction } from "../../types/transactions";
import { dataSummary } from "../../components/UI/assets";
import axiosInstance from "../../api/axios";
import {
  CustomBarChartData,
  CustomRadarChartData,
  CustomPieChartData,
  CustomPieChartColors,
  CustomLineChartData,
} from "../../components/UI/assets";
import SummaryAvgScore from "../../components/UI/summaryCards/SummaryAvgScore";
import SummaryAvgTime from "../../components/UI/summaryCards/SummaryAvgTime";
import SummaryTotal from "../../components/UI/summaryCards/SummaryTotal";
const Dashboard = ({ pageTitle }: { pageTitle: string }) => {
  // const [transactions, setTransactions] = useState([]);
  const [transactions, setTransactions] = useState<ITransaction[]>([]);

  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await axiosInstance.get("/api/transactions/");
        setTransactions(response.data);
      } catch (error) {
        console.log(error);
      }
    };

    fetchTransactions();
  }, []);

  return (
    <div className="grid grid-cols-1 gap-[19px] md:grid-cols-3 ">
      {/* Sales Summary Card */}

      <SummaryAvgScore
        className=" col-span-3 md:col-span-1"
        transactions={transactions}
      />
      {/* Customers Summary Card */}
      <SummaryAvgTime
        className="col-span-3 md:col-span-1"
        transactions={transactions}
      />

      {/* Orders Summary Card */}
      <SummaryTotal
        className="col-span-3 md:col-span-1"
        transactions={transactions}
      />

      {/* Charts and Additional Cards */}
      {/* Adjusting for a 2-column layout on md screens and using available cols on larger screens */}

      {/* <div className="grid grid-cols-3"> */}
      <CustomLineChart
        className="col-span-3 lg:col-span-2"
        data={CustomLineChartData}
      />

      <CustomBarChart
        className="col-span-3 lg:col-span-1"
        data={CustomBarChartData}
      />
      {/* </div> */}

      {/* <CustomPieChart
        className=" md:col-span-1 "
        // data={CustomRadarChartData}
        data={CustomPieChartData}
        colors={CustomPieChartColors}
      /> */}

      {/* Recent Orders
      <div className="md:col-span-2 lg:col-span-3 xl:col-span-1">
        <RecentOrders />
      </div> */}
    </div>
  );
};

export default Dashboard;
