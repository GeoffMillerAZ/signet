import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import DashboardShell from "@/components/DashboardShell";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Signet Console",
  description: "Autonomous Governance Dashboard",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="h-full bg-gray-950">
      <body className={`${inter.className} h-full`}>
        <DashboardShell>
          {children}
        </DashboardShell>
      </body>
    </html>
  );
}