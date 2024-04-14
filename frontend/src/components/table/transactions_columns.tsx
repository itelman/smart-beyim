"use client";

import { ColumnDef } from "@tanstack/react-table";

import { Checkbox } from "../shadcn/ui/checkbox";

import { Task } from "./data/schema";
import { DataTableColumnHeader } from "./data-table-column-header";
import { DataTableRowActions } from "./data-table-row-actions";

export const columns: ColumnDef<Task>[] = [
  {
    accessorKey: "document_id",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="№ Документа " />
    ),
    cell: ({ row }) => (
      <div className="w-[120px]">{row.getValue("document_id")}</div>
    ),
    enableSorting: false,
    enableHiding: false,
  },

  {
    accessorKey: "operation_date",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Дата операции" />
    ),
    cell: ({ row }) => (
      <div className="w-[120px]">{row.getValue("operation_date")}</div>
    ),
  },
  {
    accessorKey: "debet",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Дебет" />
    ),
    cell: ({ row }) => <div className="w-[120px]">{row.getValue("debet")}</div>,
  },
  {
    accessorKey: "credit",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Кредит" />
    ),
    cell: ({ row }) => (
      <div className="w-[120px]">{row.getValue("credit")}</div>
    ),
  },
  {
    accessorKey: "beneficiary_name",
    header: ({ column }) => (
      <DataTableColumnHeader
        column={column}
        title="Наименование бенефициара/отправителя денег"
      />
    ),
    cell: ({ row }) => (
      <div className="w-[250px]">{row.getValue("beneficiary_name")}</div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "beneficiary_bank_bic",
    header: ({ column }) => (
      <DataTableColumnHeader
        column={column}
        title="БИК банка бенефициара (отправителя денег)"
      />
    ),
    cell: ({ row }) => (
      <div className="w-[250px]">{row.getValue("beneficiary_bank_bic")}</div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "beneficiary_iic",
    header: ({ column }) => (
      <DataTableColumnHeader
        column={column}
        title="ИИК бенефициара/отправителя денег"
      />
    ),
    cell: ({ row }) => (
      <div className="w-[250px]">{row.getValue("beneficiary_iic")}</div>
    ),
    enableSorting: false,
  },
  {
    accessorKey: "knp",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="КНП" />
    ),
    cell: ({ row }) => <div className="w-[120px]">{row.getValue("knp")}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "payment_purpose",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Назначение платежа" />
    ),
    cell: ({ row }) => (
      <div className="w-[250px]">{row.getValue("payment_purpose")}</div>
    ),
    enableSorting: false,
  },

  // {
  //   id: "actions",
  //   cell: ({ row }) => <DataTableRowActions row={row} />,
  // },
];
