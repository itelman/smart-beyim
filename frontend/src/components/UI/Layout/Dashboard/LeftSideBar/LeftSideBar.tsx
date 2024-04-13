import DashboardNavigation from "./Navigation";

import Icon from "../../../Icon/Icon";
import Action from "./Action";

// TODO:REFACTOR LINK INTO SEPARATE LINK COMPONENT WITH VARIANTS

const LeftSideBar = () => {
  return (
    <div className="leftsidebar group">
      <div className="flex flex-col gap-[62px]">
        <div className="ml-[2.5px] flex items-center gap-[3px]">
          <Icon name="LogoSvg" width={30} height={30} />
          <span className=" hidden text-center font-headings text-subHeading3 font-bold text-[#45464E]  2xl:block">
            Beyim Analytics
          </span>
        </div>
        <DashboardNavigation />
      </div>

      <Action />
    </div>
  );
};

export default LeftSideBar;
