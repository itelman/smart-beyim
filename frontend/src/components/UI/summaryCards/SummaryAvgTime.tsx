import React from "react";
import { dataSummary } from "../assets";
import { ITransaction } from "../../../types/transactions";
import SummaryCard from "./SummaryCard";
import { useState } from "react";
import { getDateRange } from "../../../utils/utils";
interface ISummaryCardProps extends React.HTMLAttributes<HTMLDivElement> {
  transactions: ITransaction[];
}

const SummaryAvgTime: React.FC<ISummaryCardProps> = ({
  transactions,
  className,
}) => {
  //  intitial state for the filterDateRange is last 7 days, it should be a today date minus 7 days
  const today = new Date();

  const [dateFrom, setDateFrom] = useState<string>("thisWeek");

  const handleDateChange = (value: string) => {
    setDateFrom(value);
  };

  const [startDate, endDate] = getDateRange(dateFrom);

  const filterTransactions = transactions.filter((transaction) => {
    return (
      new Date(transaction.date).getTime() > startDate.getTime() &&
      new Date(transaction.date).getTime() < endDate.getTime()
    );
  });

  const totalSales = filterTransactions.reduce((acc, transaction) => {
    return acc + transaction.credit;
  }, 0);

  const transactionsVolume = filterTransactions.length;

  const salesSummary = {
    ...dataSummary.salesSummary,
    data: [
      { label: "Average time spent", value: totalSales, valueCurrency: "тг" },
    ],
  };

  return (
    <SummaryCard
      iconName={salesSummary.iconName}
      onDateChange={handleDateChange}
      dateFrom={dateFrom}
      data={salesSummary.data}
      highlight={salesSummary.highlight}
      className={className}
    />
  );
};

export default SummaryAvgTime;
