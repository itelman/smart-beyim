import { useEffect, useState } from "react";
import Icon from "../../Icon/Icon";
import { useLocation } from "react-router-dom";

const TopNav = () => {
  const [pageTitle, setPageTitle] = useState<string>("");

  const location = useLocation();

  useEffect(() => {
    setPageTitle(() => document.title);
  }, [location.pathname]);

  return (
    <>
      <div className="topbar_nav">
        <h6 className="font-headings text-subHeading3 font-medium">
          {pageTitle}
        </h6>
        <div className="flex gap-5">
          <div className="flex items-center gap-[20px]">
            <div className="flex items-center gap-[20px] rounded-[8px] border-none bg-secondary-200 px-[6px] py-[12px] sm:px-[12px] sm:py-[12px]">
              <p className="my-[1.5px] hidden text-p2 sm:block  ">
                Nanny’s Shop
              </p>
              <Icon
                name="ChevronDownSvg"
                className="mx-[5px] my-[7.5] cursor-pointer"
                width="12"
                height="8"
              />
            </div>
            <Icon
              name="NotificationSvg"
              className="cursor-pointer"
              width="20"
              height="20"
            />
            <img
              src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"
              width="32px"
              height="32px"
              className="cursor-pointer rounded-[8px] "
              alt=""
            />
          </div>
        </div>
      </div>
    </>
  );
};

export default TopNav;
