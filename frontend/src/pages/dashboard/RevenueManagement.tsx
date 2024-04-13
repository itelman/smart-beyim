import { useEffect, useState } from "react";
import { dataSummary } from "../../components/UI/assets";
import { Table } from "../../components/shadcn/ui/table";
// import { Payment, columns } from "../../components/table/columnsPrev";
import { columns } from "../../components/table/transactions_columns";

// import { DataTable } from "../../components/table/DataTable";
import { DataTable } from "../../components/table/data-table";
import data from "../../components/table/data/transactions.json";
import { z } from "zod";
import { taskSchema } from "../../components/table/data/schema";

const RevenueManagement = ({ pageTitle }: { pageTitle: string }) => {
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);

  return (
    <div className="grid grid-cols-1 gap-[19px] md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-dashboard-layout">
      <DataTable
        className="relative col-span-1 overflow-hidden md:col-span-2 lg:col-span-3"
        data={data}
        columns={columns}
      />
    </div>
  );
};

export default RevenueManagement;
