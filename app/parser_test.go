package main

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

var site = `
<!DOCTYPE html>
<html lang="en-US" dir="ltr">

<head>
    <title>Paradise Lost - Obsidian (Limited Edition) (2020) - Currently Leaked - Kingdom Leaks</title>
    <link rel="apple-touch-icon" sizes="76x76" href="/apple-touch-icon.png" />
    <link rel="manifest" href="/site.webmanifest" />
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#577a9c" />
    <meta name="msapplication-TileColor" content="#da532c" />
    <meta name="theme-color" content="#ffffff" />
    <!--[if lt IE 9]>
            <link rel="stylesheet" type="text/css" href="https://kingdom-leaks.com/uploads/css_built_6/5e61784858ad3c11f00b5706d12afe52_ie8.css.31abf9a824d94f08aa15bf0673fff832.css">
        <script src="//kingdom-leaks.com/applications/core/interface/html5shiv/html5shiv.js"></script>
    <![endif]-->
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="twitter:card" content="summary" />
    <meta name="description" content="Leaked Release - May 15th, 2020 Genre - Gothic Metal, Doom Metal Quality - MP3, 160 kbps CBR Tracklist: 1. Darker Thoughts (05:46) 2. Fall from Grace (05:42) 3. Ghosts (04:35) 4. The Devil Embraced
(06:08) 5. Forsaken (04:30) 6. Serenity (04:46) 7. Ending Days (04:36) 8. Hope Dies Young (04:02) 9..." />
    <meta property="og:title" content="Paradise Lost - Obsidian (Limited Edition) (2020)" />
    <meta property="og:type" content="object" />
    <meta property="og:url"
        content="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/" />
    <meta property="og:description" content="Leaked Release - May 15th, 2020 Genre - Gothic Metal, Doom Metal Quality - MP3, 160 kbps CBR Tracklist: 1. Darker Thoughts (05:46) 2. Fall from Grace (05:42) 3. Ghosts (04:35) 4. The Devil E
mbraced (06:08) 5. Forsaken (04:30) 6. Serenity (04:46) 7. Ending Days (04:36) 8. Hope Dies Young (04:02) 9..." />
    <meta property="og:updated_time" content="2020-05-04T17:00:56Z" />
    <meta property="og:site_name" content="Kingdom Leaks" />
    <meta property="og:locale" content="en_US" />
    <meta name="theme-color" content="#292929" />
    <link rel="canonical"
        href="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/" />
    <link rel="alternate" type="application/rss+xml" title="The Kingdom Leaks Everything Feed"
        href="https://kingdom-leaks.com/index.php?/rss/3-the-kingdom-leaks-everything-feed.xml/" />
    <link rel="alternate" type="application/rss+xml" title="New Downloads Only Feed" h
        ref="https://kingdom-leaks.com/index.php?/rss/1-new-downloads-only-feed.xml/" />
    <link rel="alternate" type="application/rss+xml" title="New Downloads Only (without Singles) Feed" href="https://kingdom-leaks.com/index.php?/rss/5-new-downloads-only-without-singles-f
eed.xml/" />
    <link rel="alternate" type="application/rss+xml" title="Newsroom Feed"
        href="https://kingdom-leaks.com/index.php?/rss/2-newsroom-feed.xml/" />
    <script src="https://ajax.googleapis.com/ajax/libs/webfont/1.6.26/webfont.js" defer=""></script>
    <script defer="">
        window.addEventListener('DOMContentLoaded', (event) => {
            WebFont.load({
                google: {
                    families: ["Roboto:300,300i,400,400i,700,700i"]
                }
            });
        });
    </script>
    <!--<link href="https://fonts.googleapis.com/css?family=Roboto:300,300i,400,400i,700,700i" rel="stylesheet" referrerpolicy="origin">-->
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/341e4a57816af3ba440d891ca87450ff_framework.css.fa367e93e11eba149c8e9a674bd86b05.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/05e81b71abe4f22d6eb8d1a929494829_responsive.css.81851d5a9c79e8f1252bacd5dd93a694.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/ec0c06d47f161faa24112e8cbf0665bc_chatbox.css.b40ff2332ae1f98d35c2a049e09786b0.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/90eb5adf50a8c640f633d47fd7eb1778_core.css.81ecf6dcc29d650fbef1a22577324013.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/5a0da001ccc2200dc5625c3f3934497d_core_responsive.css.997a6a7b099aeeade4a621f845d4c7a6.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/62e269ced0fdab7e30e026f1d30ae516_forums.css.fa696074a2639db7bcd325e986ca93e8.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/76e62c573090645fb99a15a363d8620e_forums_responsive.css.f62bbc45b3f8dc8055e810db7cabe7ba.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/c8cd9abd157846b1207e9b2320977b5a_slider.css.068b8d427ec1faee59ce6d1bc2ecd89c.css?v=026d4e282a"
        media="all" />
    <link rel="stylesheet"
        href="https://kingdom-leaks.com/uploads/css_built_6/258adbb6e4f3e83cd3b355f84e3fa002_custom.css.a9f911595f7271b11ec7afc2a7bd293f.css?v=026d4e282a"
        media="all" />
    <!--START MOBILE TOGGLE-->
    <script>
        window.addEventListener('DOMContentLoaded', (event) => {
            $(".toggle_button_left").click(function () {
                $("#ipsLayout_sidebar_left").slideToggle();
            });
        });
    </script>
    <!--END MOBILE TOGGLE-->
    <!--START CUSTOM JS SCRIPTS-->
    <!--<script type="text/javascript" src="/js/iframeResizer.min.js?v=1.0"></script>-->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/iframe-resizer/4.2.10/iframeResizer.min.js"
        integrity="sha256-0FsDr6k3iiIaao/F1olkJHUfEU/eGSYClQ7ZhVc2md8=" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.0/jquery.min.js"
        integrity="sha256-xNzN2a4ltkB44Mc/Jz3pT4iU1cmeR0FkXs4pru/JxaQ=" crossorigin="anonymous"></script>
    <!--<script>
            window.lazySizesConfig = window.lazySizesConfig || {};
            window.lazySizesConfig.init = false;
window.addEventListener('DOMContentLoaded', (event) => {
      lazySizes.init();
  console.log("lazy loaded.");
});
</script>-->
    <script src="/js/lazysizes.min.js" async=""></script>
    <script src="/js/ls.unveilhooks.min.js" async=""></script>
    <script src="/js/lastfm.js?v=1.1.10" defer=""></script>
    <!--END CUSTOM JS SCRIPTS-->
    <!--START CHRISTMAS DECORATIONS-->
    <!--<script src="/js/snowstorm.js"></script>
<script>
snowStorm.vMaxX = 4;
snowStorm.vMaxY = 0.01;
</script>-->
    <!--END CHRISTMAS DECORATIONS-->
</head>

<body class="ipsApp ipsApp_front ipsJS_none ipsClearfix " data-controller="core.front.core.app,plugins.minimizequote"
    data-message="" data-pageapp="forums" data-pagelocation="front" data-pagemodule="forums" data-pagecontroller="topic"
    itemscope="" itemtyp e="http://schema.org/WebSite">
    <meta itemprop="url" content="https://kingdom-leaks.com/" />
    <a href="#elContent" class="ipsHide" title="Go to main content on this page" accesskey="m">Jump to content</a>
    <!--  Header Slider -->
    <div id="ipsLayout_header" class="ipsClearfix ipsLayout_container v-header">
        <ul id="elMobileNav" class="ipsList_inline ipsResponsive_hideDesktop ipsResponsive_block"
            data-controller="core.front.core.mobileNav" data-default="forums_topic">
            <li id="elMobileBreadcrumb">
                <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/">
                    <span>Currently Leaked</span>
                </a>
            </li>
            <!--
<li >
    <a data-action="defaultStream" class='ipsType_light'  href='https://kingdom-leaks.com/index.php?/discover/15/'><i class='icon-newspaper'></i></a>
</li>-->
            <li class="ipsJS_show">
                <a
                    href="https://kingdom-leaks.com/index.php?/search/&amp;type=forums_topic&amp;search_in=titles&amp;sortby=newest"><i
                        class="fa fa-search"></i></a>
            </li>
            <li data-ipsdrawer="" data-ipsdrawer-drawerelem="#elMobileDrawer">
                <a href="#">
                    <i class="fa fa-navicon"></i>
                </a>
            </li>
        </ul>
        <header>
            <div class="vn-gh">
                <div class="v-logo">
                    <a href="https://kingdom-leaks.com/" id="elLogo" accesskey="1"><img
                            src="https://kingdom-leaks.com/uploads/monthly_2020_02/KLll.png.577e903d7902e67c46afa23ca0fa6050.png"
                            alt="Kingdom Leaks" /></a>
                </div>
                <!-- Search Input-->
                <div id="elSearch" style="margin-top:15px;" class="" data-controller="core.front.core.quickSearch">
                    <form accept-charset="utf-8"
                        action="//kingdom-leaks.com/index.php?/search/&amp;search_and_or=and&amp;search_in=titles&amp;sortby=newest"
                        method="post">
                        <input type="search" id="elSearchField" placeholder="Search..." name="q" />
                        <button class="cSearchSubmit" type="submit"><i class="fa fa-search"></i></button>
                        <div id="elSearchExpanded" style="display:none">
                            <div class="ipsMenu_title">
                                Search In
                            </div>
                            <ul class="ipsSideMenu_list ipsSideMenu_withRadios ipsSideMenu_small ipsType_normal"
                                data-ipssidemenu="" data-ipssidemenu-type="radio" data-ipssidemenu-responsive="false"
                                data-role="searchContexts">
                                <li>
                                    <span class="ipsSideMenu_item ipsSideMenu_itemActive" data-ipsmenuvalue="all">
                                        <input type="radio" name="type" value="all" checked=""
                                            id="elQuickSearchRadio_type_all" />
                                        <label for="elQuickSearchRadio_type_all"
                                            id="elQuickSearchRadio_type_all_label">Everywhere</label>
                                    </span>
                                </li>
                                <li>
                                    <span class="ipsSideMenu_item" data-ipsmenuvalue="forums_topic">
                                        <input type="radio" name="type" value="forums_topic"
                                            id="elQuickSearchRadio_type_forums_topic" />
                                        <label for="elQuickSearchRadio_type_forums_topic"
                                            id="elQuickSearchRadio_type_forums_topic_label">Topics</label>
                                    </span>
                                </li>
                                <li>
                                    <span class="ipsSideMenu_item"
                                        data-ipsmenuvalue="contextual_{&#34;type&#34;:&#34;forums_topic&#34;,&#34;nodes&#34;:109}">
                                        <input type="radio" name="type"
                                            value="contextual_{&#34;type&#34;:&#34;forums_topic&#34;,&#34;nodes&#34;:109}"
                                            id="elQuickSearchRadio_type_contextual_{&#34;type&#34;:&#34;forums_topic
&#34;,&#34;nodes&#34;:109}" />
                                        <label
                                            for="elQuickSearchRadio_type_contextual_{&#34;type&#34;:&#34;forums_topic&#34;,&#34;nodes&#34;:109}"
                                            id="elQuickSearchRadio_type_contextual_{&#34;type&#34;:&#34;forums_topic&#3
4;,&#34;nodes&#34;:109}_label">This Forum</label>
                                    </span>
                                </li>
                                <li>
                                    <span class="ipsSideMenu_item"
                                        data-ipsmenuvalue="contextual_{&#34;type&#34;:&#34;forums_topic&#34;,&#34;item&#34;:37899}">
                                        <input type="radio" name="type"
                                            value="contextual_{&#34;type&#34;:&#34;forums_topic&#34;,&#34;item&#34;:37899}"
                                            id="elQuickSearchRadio_type_contextual_{&#34;type&#34;:&#34;forums_topi
c&#34;,&#34;item&#34;:37899}" />
                                        <label
                                            for="elQuickSearchRadio_type_contextual_{&#34;type&#34;:&#34;forums_topic&#34;,&#34;item&#34;:37899}"
                                            id="elQuickSearchRadio_type_contextual_{&#34;type&#34;:&#34;forums_topic&#
34;,&#34;item&#34;:37899}_label">This Topic</label>
                                    </span>
                                </li>
                                <li data-role="showMoreSearchContexts">
                                    <span class="ipsSideMenu_item" data-action="showMoreSearchContexts"
                                        data-exclude="forums_topic">
                                        More options...
                                    </span>
                                </li>
                            </ul>
                            <div class="ipsMenu_title">
                                Find results that contain...
                            </div>
                            <ul class="ipsSideMenu_list ipsSideMenu_withRadios ipsSideMenu_small ipsType_normal"
                                role="radiogroup" data-ipssidemenu="" data-ipssidemenu-type="radio"
                                data-ipssidemenu-responsive="false" data-filtertype="andOr">
                                <li>
                                    <span class="ipsSideMenu_item ipsSideMenu_itemActive" data-ipsmenuvalue="and">
                                        <input type="radio" name="search_and_or" value="and" checked=""
                                            id="elRadio_andOr_and" />
                                        <label for="elRadio_andOr_and" id="elField_andOr_label_and"><em>All</em> of my
                                            search term words</label>
                                    </span>
                                </li>
                                <li>
                                    <span class="ipsSideMenu_item " data-ipsmenuvalue="or">
                                        <input type="radio" name="search_and_or" value="or" id="elRadio_andOr_or" />
                                        <label for="elRadio_andOr_or" id="elField_andOr_label_or"><em>Any</em> of my
                                            search term words</label>
                                    </span>
                                </li>
                            </ul>
                            <div class="ipsMenu_title">
                                Find results in...
                            </div>
                            <ul class="ipsSideMenu_list ipsSideMenu_withRadios ipsSideMenu_small ipsType_normal"
                                role="radiogroup" data-ipssidemenu="" data-ipssidemenu-type="radio"
                                data-ipssidemenu-responsive="false" data-filtertype="searchIn">
                                <li>
                                    <span class="ipsSideMenu_item ipsSideMenu_itemActive" data-ipsmenuvalue="all">
                                        <input type="radio" name="search_in" value="all" checked=""
                                            id="elRadio_searchIn_and" />
                                        <label for="elRadio_searchIn_and" id="elField_searchIn_label_all">Content titles
                                            and body</label>
                                    </span>
                                </li>
                                <li>
                                    <span class="ipsSideMenu_item" data-ipsmenuvalue="titles">
                                        <input type="radio" name="search_in" value="titles"
                                            id="elRadio_searchIn_titles" />
                                        <label for="elRadio_searchIn_titles" id="elField_searchIn_label_titles">Content
                                            titles only</label>
                                    </span>
                                </li>
                            </ul>
                        </div>
                    </form>
                </div>
                <ul id="elUserNav" style="margin-top:22px"
                    class="ipsList_inline cSignedOut ipsClearfix ipsResponsive_hidePhone ipsResponsive_block">
                    <li id="elSignInLink">
                        <a href="//kingdom-leaks.com/index.php?/login/" data-ipsmenu-closeonclick="false"
                            data-ipsmenu="" id="elUserSignIn">
                            Existing user? Sign In  <i class="fa fa-caret-down"></i>
                        </a>
                        <div id="elUserSignIn_menu" class="ipsMenu ipsMenu_auto ipsHide">
                            <form accept-charset="utf-8" method="post" action="//kingdom-leaks.com/index.php?/login/"
                                data-controller="core.global.core.login">
                                <input type="hidden" name="csrfKey" value="538d6993577ac1a4f5e0fc418dc51e57" />
                                <input type="hidden" name="ref"
                                    value="aHR0cHM6Ly9raW5nZG9tLWxlYWtzLmNvbS9pbmRleC5waHA/L2ZvcnVtcy90b3BpYy8zNzg5OS1wYXJhZGlzZS1sb3N0LW9ic2lkaWFuLWxpbWl0ZWQtZWRpdGlvbi0yMDIwLw==" />
                                <div data-role="loginForm">
                                    <div class="ipsPad ipsForm ipsForm_vertical">
                                        <h4 class="ipsType_sectionHead">Sign In</h4>
                                        <br /><br />
                                        <ul class="ipsList_reset">
                                            <li class="ipsFieldRow ipsFieldRow_noLabel ipsFieldRow_fullWidth">
                                                <input type="text" placeholder="Display Name" name="auth" />
                                            </li>
                                            <li class="ipsFieldRow ipsFieldRow_noLabel ipsFieldRow_fullWidth">
                                                <input type="password" placeholder="Password" name="password" />
                                            </li>
                                            <li class="ipsFieldRow ipsFieldRow_checkbox ipsClearfix">
                                                <span class="ipsCustomInput">
                                                    <input type="checkbox" name="remember_me" id="remember_me_checkbox"
                                                        value="1" checked="" aria-checked="true" />
                                                    <span></span>
                                                </span>
                                                <div class="ipsFieldRow_content">
                                                    <label class="ipsFieldRow_label" for="remember_me_checkbox">Remember
                                                        me</label>
                                                    <span class="ipsFieldRow_desc">Not recommended on shared
                                                        computers</span>
                                                </div>
                                            </li>
                                            <li class="ipsFieldRow ipsFieldRow_fullWidth">
                                                <br />
                                                <button type="submit" name="_processLogin" value="usernamepassword"
                                                    class="ipsButton ipsButton_primary ipsButton_small"
                                                    id="elSignIn_submit">Sign In</button>
                                                <br />
                                                <p class="ipsType_right ipsType_small">
                                                    <a href="https://kingdom-leaks.com/index.php?/lostpassword/"
                                                        data-ipsdialog="" data-ipsdialog-title="Forgot your password?">
                                                        Forgot your password?</a>
                                                </p>
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </li>
                    <li>
                        <a href="https://kingdom-leaks.com/index.php?/register/" id="elRegisterButton"
                            class="ipsButton ipsButton_normal ipsButton_primary">
                            Sign Up
                        </a>
                    </li>
                </ul>
            </div>
        </header>
        <!--
                <nav data-controller='core.front.core.navBar'>
<div class="ipsLayout_container">
    <div class='ipsNavBar_primary  ipsClearfix v-styling-d'>
                    <ul data-role="primaryNavBar" class='ipsResponsive_showDesktop ipsResponsive_block'>
                            <li aria-haspopup="true"  id='elNavSecondary_35' data-role="navBarItem" data-navApp="core" data-navExt="Menu" data-navTitle="Browse">
                                            <a href="#" data-navItem-id="35" >
                            Browse
                    </a>
                                        <ul class='ipsNavBar_secondary ipsHide'>
                            <li aria-haspopup="true"  id='elNavSecondary_8' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Forums">
                                            <a href="https://kingdom-leaks.com/index.php?/forums/" data-navItem-id="8" >
                            Forums
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_68' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="Latest Downloads">
                                            <a href="https://kingdom-leaks.com/index.php?/new/" data-navItem-id="68" >
                            Latest Downloads
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_110' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Currently Leaked">
                                            <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/" data-navItem-id="110" >
                            Currently Leaked
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_109' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="Latest Singles">
                                            <a href="https://kingdom-leaks.com/index.php?/singles/" data-navItem-id="109" >
                            Latest Singles
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_108' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="Trending Leaks">
                                            <a href="https://kingdom-leaks.com/index.php?/trending/" data-navItem-id="108" >
                            Trending Leaks
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_74' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="Latest Quality Updates">
                                            <a href="https://kingdom-leaks.com/index.php?/hq/" data-navItem-id="74" >
                            Latest Quality Updates
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_41' data-role="navBarItem" data-navApp="core" data-navExt="Leaderboard" data-navTitle="Leaderboard">
                                            <a href="https://kingdom-leaks.com/index.php?/pastleaders/" data-navItem-id="41" >
                            Leaderboard
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_42' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Online Users">
                                            <a href="https://kingdom-leaks.com/index.php?/online/&filter=filter_loggedin" data-navItem-id="42" >
                            Online Users
                    </a>
                    </li>
                                                <li class='ipsHide' id='elNavigationMore_35' data-role='navMore'>
                                    <a href='#' data-ipsMenu data-ipsMenu-appendTo='#elNavigationMore_35' id='elNavigationMore_35_dropdown'>More <i class='fa fa-caret-down'></i></a>
                                    <ul class='ipsHide ipsMenu ipsMenu_auto' id='elNavigationMore_35_dropdown_menu' data-role='moreDropdown'></ul>
                            </li>
                    </ul>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_81' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Submit a Leak">
                                            <a href="https://kingdom-leaks.com/index.php?/forms/2-submit-a-leak/" data-navItem-id="81" >
                            Submit a Leak
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_43' data-role="navBarItem" data-navApp="core" data-navExt="Menu" data-navTitle="About Us">
                                            <a href="#" data-navItem-id="43" >
                            About Us
                    </a>
                                        <ul class='ipsNavBar_secondary ipsHide'>
                            <li aria-haspopup="true"  id='elNavSecondary_104' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Staff & Trusted">
                                            <a href="https://kingdom-leaks.com/index.php?/staff/" data-navItem-id="104" >
                            Staff & Trusted
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_69' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="Site Rules">
                                            <a href="https://kingdom-leaks.com/index.php?/rules/" data-navItem-id="69" >
                            Site Rules
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_105' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="How to Download">
                                            <a href="https://kingdom-leaks.com/index.php?/forums/topic/28592-how-to-download-from-kingdom-leaks/" data-navItem-id="105" >
                            How to Download
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_70' data-role="navBarItem" data-navApp="core" data-navExt="Promoted" data-navTitle="Our Picks">
                                            <a href="https://kingdom-leaks.com/index.php?/ourpicks/" data-navItem-id="70" >
                            Our Picks
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_84' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Donate">
                                            <a href="https://kingdom-leaks.com/index.php?/donations/" data-navItem-id="84" >
                            Donate
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_77' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="User Classes">
                                            <a href="https://kingdom-leaks.com/index.php?/user-classes/" data-navItem-id="77" >
                            User Classes
                    </a>
                    </li>
                                                <li class='ipsHide' id='elNavigationMore_43' data-role='navMore'>
                                    <a href='#' data-ipsMenu data-ipsMenu-appendTo='#elNavigationMore_43' id='elNavigationMore_43_dropdown'>More <i class='fa fa-caret-down'></i></a>
                                    <ul class='ipsHide ipsMenu ipsMenu_auto' id='elNavigationMore_43_dropdown_menu' data-role='moreDropdown'></ul>
                            </li>
                    </ul>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_65' data-role="navBarItem" data-navApp="core" data-navExt="Menu" data-navTitle="Misc">
                                                    <a href="#" id="elNavigation_65" data-ipsMenu data-ipsMenu-appendTo='#elNavSecondary_65' data-ipsMenu-activeClass='ipsNavActive_menu' data-navItem-id="65" >
                            Misc
                    </a>
                                    <ul id="elNavigation_65_menu" class="ipsMenu ipsMenu_auto ipsHide">
                                                    <li class='ipsMenu_item' >
                    <a href='https://kingdom-leaks.com/index.php?/forums/forum/54-requests/' >
                            Request
                    </a>
            </li>
                                            </ul>
                                <ul class='ipsNavBar_secondary ipsHide'>
                    <li aria-haspopup="true"  id='elNavSecondary_94' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Request">
                                    <a href="https://kingdom-leaks.com/index.php?/forums/forum/54-requests/" data-navItem-id="94" >
                            Request
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_57' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Artwork Finder">
                                            <a href="https://kingdom-leaks.com/index.php?/artwork/" data-navItem-id="57" >
                            Artwork Finder
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_64' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Chat">
                                            <a href="https://kingdom-leaks.com/index.php?/chat/" data-navItem-id="64" >
                            Chat
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_60' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="Release Notes">
                                            <a href="https://kingdom-leaks.com/index.php?/release-notes/" data-navItem-id="60" >
                            Release Notes
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_66' data-role="navBarItem" data-navApp="cms" data-navExt="Pages" data-navTitle="Feature Plan">
                                            <a href="https://kingdom-leaks.com/index.php?/feature-plan/" data-navItem-id="66" >
                            Feature Plan
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_49' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="KLStatus">
                                            <a href="https://status.kingdom-leaks.com/" data-navItem-id="49" >
                            KLStatus
                    </a>
                    </li>
                        <li aria-haspopup="true"  id='elNavSecondary_50' data-role="navBarItem" data-navApp="core" data-navExt="CustomItem" data-navTitle="#SaveKL">
                                            <a href="https://web.archive.org/save/kingdom-leaks.com" data-navItem-id="50" >
                            #SaveKL
                    </a>
                    </li>
                                                <li class='ipsHide' id='elNavigationMore_65' data-role='navMore'>
                                    <a href='#' data-ipsMenu data-ipsMenu-appendTo='#elNavigationMore_65' id='elNavigationMore_65_dropdown'>More <i class='fa fa-caret-down'></i></a>
                                    <ul class='ipsHide ipsMenu ipsMenu_auto' id='elNavigationMore_65_dropdown_menu' data-role='moreDropdown'></ul>
                            </li>
                    </ul>
                    </li>
                                                <li class='ipsHide' id='elNavigationMore' data-role='navMore'>
                                    <a href='#' data-ipsMenu data-ipsMenu-appendTo='#elNavigationMore' id='elNavigationMore_dropdown'>More</a>
                                    <ul class='ipsNavBar_secondary ipsHide' data-role='secondaryNavBar'>
                                            <li class='ipsHide veilonHide' id='elNavigationMore_more' data-role='navMore'>
                                                    <a href='#' data-ipsMenu data-ipsMenu-appendTo='#elNavigationMore_more' id='elNavigationMore_more_dropdown'>More <i class='fa fa-caret-down'></i></a>
                                                    <ul class='ipsHide ipsMenu ipsMenu_auto' id='elNavigationMore_more_dropdown_menu' data-role='moreDropdown'></ul>
                                            </li>
                                    </ul>
                            </li>
                    </ul>
                                    </div>
                    <ul class="veilon-social-icons ipsList_inline ipsPos_right">
                                    <a title="Facebook" target="_blank" href="https://www.facebook.com/KingdomLeaks/
" data-ipsTooltip><i class="fa fa-facebook ipsType_larger"></i></a>
            <a title="Twitter" target="_blank" href="https://twitter.com/masterleakers
" data-ipsTooltip><i class="fa fa-twitter ipsType_larger"></i></a>
            <a title="Instagram" target="_blank" href="https://www.instagram.com/kingdomleaks/
" data-ipsTooltip><i class="fa fa-instagram ipsType_larger"></i></a>
            <a title="Spotify" target="_blank" href="https://open.spotify.com/user/p051x03x8hqzylgwgsg8ayt4r
" data-ipsTooltip><i class="fa fa-spotify ipsType_larger"></i></a>
            <a title="Apple Music" target="_blank" href="https://music.apple.com/profile/kingdomleaks
" data-ipsTooltip><i class="fa fa-apple ipsType_larger"></i></a>
            <a title="YouTube" target="_blank" href="https://www.youtube.com/channel/UCxm0Vz6bzNFu5357QNQ5lyA" data-ipsTooltip><i class="fa fa-youtube-play ipsType_larger"></i></a>
</ul>
<ul id="elUserNav" style="margin-top:22px" class="ipsList_inline cSignedOut ipsClearfix ipsResponsive_hidePhone ipsResponsive_block">
    <li id="elSignInLink">
            <a href="//kingdom-leaks.com/index.php?/login/" data-ipsmenu-closeonclick="false" data-ipsmenu id="elUserSignIn">
                    Existing user? Sign In  <i class="fa fa-caret-down"></i>
            </a>
            <div id='elUserSignIn_menu' class='ipsMenu ipsMenu_auto ipsHide'>
<form accept-charset='utf-8' method='post' action='//kingdom-leaks.com/index.php?/login/' data-controller="core.global.core.login">
    <input type="hidden" name="csrfKey" value="538d6993577ac1a4f5e0fc418dc51e57">
    <input type="hidden" name="ref" value="aHR0cHM6Ly9raW5nZG9tLWxlYWtzLmNvbS9pbmRleC5waHA/L2ZvcnVtcy90b3BpYy8zNzg5OS1wYXJhZGlzZS1sb3N0LW9ic2lkaWFuLWxpbWl0ZWQtZWRpdGlvbi0yMDIwLw==">
    <div data-role="loginForm">
        <div class="ipsPad ipsForm ipsForm_vertical">
<h4 class="ipsType_sectionHead">Sign In</h4>
<br><br>
<ul class='ipsList_reset'>
    <li class="ipsFieldRow ipsFieldRow_noLabel ipsFieldRow_fullWidth">
                            <input type="text" placeholder="Display Name" name="auth">
        </li>
    <li class="ipsFieldRow ipsFieldRow_noLabel ipsFieldRow_fullWidth">
            <input type="password" placeholder="Password" name="password">
    </li>
    <li class="ipsFieldRow ipsFieldRow_checkbox ipsClearfix">
            <span class="ipsCustomInput">
                    <input type="checkbox" name="remember_me" id="remember_me_checkbox" value="1" checked aria-checked="true">
                    <span></span>
            </span>
            <div class="ipsFieldRow_content">
                    <label class="ipsFieldRow_label" for="remember_me_checkbox">Remember me</label>
                    <span class="ipsFieldRow_desc">Not recommended on shared computers</span>
            </div>
    </li>
        <li class="ipsFieldRow ipsFieldRow_fullWidth">
            <br>
            <button type="submit" name="_processLogin" value="usernamepassword" class="ipsButton ipsButton_primary ipsButton_small" id="elSignIn_submit">Sign In</button>
                                <br>
                    <p class="ipsType_right ipsType_small">
                                                            <a href='https://kingdom-leaks.com/index.php?/lostpassword/' data-ipsDialog data-ipsDialog-title='Forgot your password?'>
                                                    Forgot your password?</a>
                    </p>
                    </li>
</ul>
</div>
    </div>
</form>
</div>
    </li>
                <li>
                    <a href="https://kingdom-leaks.com/index.php?/register/" id="elRegisterButton" class="ipsButton ipsButton_normal ipsButton_primary">
                            Sign Up
                    </a>
            </li>
        </ul>
                            </div>
    </div>
</nav>
-->
    </div>
    <div class="v-nav">
        <nav data-controller="core.front.core.navBar">
            <div class="ipsLayout_container">
                <div class="ipsNavBar_primary  ipsClearfix v-styling-d">
                    <ul data-role="primaryNavBar" class="ipsResponsive_showDesktop ipsResponsive_block">
                        <li aria-haspopup="true" id="elNavSecondary_35" data-role="navBarItem" data-navapp="core"
                            data-navext="Menu" data-navtitle="Browse">
                            <a href="#" data-navitem-id="35">
                                Browse
                            </a>
                            <ul class="ipsNavBar_secondary ipsHide">
                                <li aria-haspopup="true" id="elNavSecondary_8" data-role="navBarItem" data-navapp="core"
                                    data-navext="CustomItem" data-navtitle="Forums">
                                    <a href="https://kingdom-leaks.com/index.php?/forums/" data-navitem-id="8">
                                        Forums
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_68" data-role="navBarItem" data-navapp="cms"
                                    data-navext="Pages" data-navtitle="Latest Downloads">
                                    <a href="https://kingdom-leaks.com/index.php?/new/" data-navitem-id="68">
                                        Latest Downloads
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_110" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Currently Leaked">
                                    <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/"
                                        data-navitem-id="110">
                                        Currently Leaked
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_109" data-role="navBarItem"
                                    data-navapp="cms" data-navext="Pages" data-navtitle="Latest Singles">
                                    <a href="https://kingdom-leaks.com/index.php?/singles/" data-navitem-id="109">
                                        Latest Singles
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_108" data-role="navBarItem"
                                    data-navapp="cms" data-navext="Pages" data-navtitle="Trending Leaks">
                                    <a href="https://kingdom-leaks.com/index.php?/trending/" data-navitem-id="108">
                                        Trending Leaks
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_74" data-role="navBarItem" data-navapp="cms"
                                    data-navext="Pages" data-navtitle="Latest Quality Updates">
                                    <a href="https://kingdom-leaks.com/index.php?/hq/" data-navitem-id="74">
                                        Latest Quality Updates
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_41" data-role="navBarItem"
                                    data-navapp="core" data-navext="Leaderboard" data-navtitle="Leaderboard">
                                    <a href="https://kingdom-leaks.com/index.php?/pastleaders/" data-navitem-id="41">
                                        Leaderboard
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_42" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Online Users">
                                    <a href="https://kingdom-leaks.com/index.php?/online/&amp;filter=filter_loggedin"
                                        data-navitem-id="42">
                                        Online Users
                                    </a>
                                </li>
                                <li class="ipsHide" id="elNavigationMore_35" data-role="navMore">
                                    <a href="#" data-ipsmenu="" data-ipsmenu-appendto="#elNavigationMore_35"
                                        id="elNavigationMore_35_dropdown">More <i class="fa fa-caret-down"></i></a>
                                    <ul class="ipsHide ipsMenu ipsMenu_auto" id="elNavigationMore_35_dropdown_menu"
                                        data-role="moreDropdown"></ul>
                                </li>
                            </ul>
                        </li>
                        <li aria-haspopup="true" id="elNavSecondary_81" data-role="navBarItem" data-navapp="core"
                            data-navext="CustomItem" data-navtitle="Submit a Leak">
                            <a href="https://kingdom-leaks.com/index.php?/forms/2-submit-a-leak/" data-navitem-id="81">
                                Submit a Leak
                            </a>
                        </li>
                        <li aria-haspopup="true" id="elNavSecondary_43" data-role="navBarItem" data-navapp="core"
                            data-navext="Menu" data-navtitle="About Us">
                            <a href="#" data-navitem-id="43">
                                About Us
                            </a>
                            <ul class="ipsNavBar_secondary ipsHide">
                                <li aria-haspopup="true" id="elNavSecondary_104" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Staff &amp; Trusted">
                                    <a href="https://kingdom-leaks.com/index.php?/staff/" data-navitem-id="104">
                                        Staff &amp; Trusted
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_69" data-role="navBarItem" data-navapp="cms"
                                    data-navext="Pages" data-navtitle="Site Rules">
                                    <a href="https://kingdom-leaks.com/index.php?/rules/" data-navitem-id="69">
                                        Site Rules
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_105" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="How to Download">
                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/28592-how-to-download-from-kingdom-leaks/"
                                        data-navitem-id="105">
                                        How to Download
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_70" data-role="navBarItem"
                                    data-navapp="core" data-navext="Promoted" data-navtitle="Our Picks">
                                    <a href="https://kingdom-leaks.com/index.php?/ourpicks/" data-navitem-id="70">
                                        Our Picks
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_84" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Donate">
                                    <a href="https://kingdom-leaks.com/index.php?/donations/" data-navitem-id="84">
                                        Donate
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_77" data-role="navBarItem" data-navapp="cms"
                                    data-navext="Pages" data-navtitle="User Classes">
                                    <a href="https://kingdom-leaks.com/index.php?/user-classes/" data-navitem-id="77">
                                        User Classes
                                    </a>
                                </li>
                                <li class="ipsHide" id="elNavigationMore_43" data-role="navMore">
                                    <a href="#" data-ipsmenu="" data-ipsmenu-appendto="#elNavigationMore_43"
                                        id="elNavigationMore_43_dropdown">More <i class="fa fa-caret-down"></i></a>
                                    <ul class="ipsHide ipsMenu ipsMenu_auto" id="elNavigationMore_43_dropdown_menu"
                                        data-role="moreDropdown"></ul>
                                </li>
                            </ul>
                        </li>
                        <li aria-haspopup="true" id="elNavSecondary_65" data-role="navBarItem" data-navapp="core"
                            data-navext="Menu" data-navtitle="Misc">
                            <a href="#" id="elNavigation_65" data-ipsmenu="" data-ipsmenu-appendto="#elNavSecondary_65"
                                data-ipsmenu-activeclass="ipsNavActive_menu" data-navitem-id="65">
                                Misc
                            </a>
                            <ul id="elNavigation_65_menu" class="ipsMenu ipsMenu_auto ipsHide">
                                <li class="ipsMenu_item">
                                    <a href="https://kingdom-leaks.com/index.php?/forums/forum/54-requests/">
                                        Request
                                    </a>
                                </li>
                            </ul>
                            <ul class="ipsNavBar_secondary ipsHide">
                                <li aria-haspopup="true" id="elNavSecondary_94" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Request">
                                    <a href="https://kingdom-leaks.com/index.php?/forums/forum/54-requests/"
                                        data-navitem-id="94">
                                        Request
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_57" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Artwork Finder">
                                    <a href="https://kingdom-leaks.com/index.php?/artwork/" data-navitem-id="57">
                                        Artwork Finder
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_64" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Chat">
                                    <a href="https://kingdom-leaks.com/index.php?/chat/" data-navitem-id="64">
                                        Chat
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_60" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="Release Notes">
                                    <a href="https://kingdom-leaks.com/index.php?/release-notes/" data-navitem-id="60">
                                        Release Notes
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_66" data-role="navBarItem" data-navapp="cms"
                                    data-navext="Pages" data-navtitle="Feature Plan">
                                    <a href="https://kingdom-leaks.com/index.php?/feature-plan/" data-navitem-id="66">
                                        Feature Plan
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_49" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="KLStatus">
                                    <a href="https://status.kingdom-leaks.com/" data-navitem-id="49">
                                        KLStatus
                                    </a>
                                </li>
                                <li aria-haspopup="true" id="elNavSecondary_50" data-role="navBarItem"
                                    data-navapp="core" data-navext="CustomItem" data-navtitle="#SaveKL">
                                    <a href="https://web.archive.org/save/kingdom-leaks.com" data-navitem-id="50">
                                        #SaveKL
                                    </a>
                                </li>
                                <li class="ipsHide" id="elNavigationMore_65" data-role="navMore">
                                    <a href="#" data-ipsmenu="" data-ipsmenu-appendto="#elNavigationMore_65"
                                        id="elNavigationMore_65_dropdown">More <i class="fa fa-caret-down"></i></a>
                                    <ul class="ipsHide ipsMenu ipsMenu_auto" id="elNavigationMore_65_dropdown_menu"
                                        data-role="moreDropdown"></ul>
                                </li>
                            </ul>
                        </li>
                        <li class="ipsHide" id="elNavigationMore" data-role="navMore">
                            <a href="#" data-ipsmenu="" data-ipsmenu-appendto="#elNavigationMore"
                                id="elNavigationMore_dropdown">More</a>
                            <ul class="ipsNavBar_secondary ipsHide" data-role="secondaryNavBar">
                                <li class="ipsHide veilonHide" id="elNavigationMore_more" data-role="navMore">
                                    <a href="#" data-ipsmenu="" data-ipsmenu-appendto="#elNavigationMore_more"
                                        id="elNavigationMore_more_dropdown">More <i class="fa fa-caret-down"></i></a>
                                    <ul class="ipsHide ipsMenu ipsMenu_auto" id="elNavigationMore_more_dropdown_menu"
                                        data-role="moreDropdown"></ul>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </div>
                <ul class="veilon-social-icons ipsList_inline ipsPos_right">
                    <a title="Facebook" target="_blank" href="https://www.facebook.com/KingdomLeaks/
" data-ipstooltip=""><i class="fa fa-facebook ipsType_larger"></i></a>
                    <a title="Twitter" target="_blank" href="https://twitter.com/masterleakers
" data-ipstooltip=""><i class="fa fa-twitter ipsType_larger"></i></a>
                    <a title="Instagram" target="_blank" href="https://www.instagram.com/kingdomleaks/
" data-ipstooltip=""><i class="fa fa-instagram ipsType_larger"></i></a>
                    <a title="Spotify" target="_blank" href="https://open.spotify.com/user/p051x03x8hqzylgwgsg8ayt4r
" data-ipstooltip=""><i class="fa fa-spotify ipsType_larger"></i></a>
                    <a title="Apple Music" target="_blank" href="https://music.apple.com/profile/kingdomleaks
" data-ipstooltip=""><i class="fa fa-apple ipsType_larger"></i></a>
                    <a title="YouTube" target="_blank" href="https://www.youtube.com/channel/UCxm0Vz6bzNFu5357QNQ5lyA"
                        data-ipstooltip=""><i class="fa fa-youtube-play ipsType_larger"></i></a>
                </ul>
                <ul id="elUserNav" style="margin-top:22px"
                    class="ipsList_inline cSignedOut ipsClearfix ipsResponsive_hidePhone ipsResponsive_block">
                    <li id="elSignInLink">
                        <a href="//kingdom-leaks.com/index.php?/login/" data-ipsmenu-closeonclick="false"
                            data-ipsmenu="" id="elUserSignIn">
                            Existing user? Sign In  <i class="fa fa-caret-down"></i>
                        </a>
                        <div id="elUserSignIn_menu" class="ipsMenu ipsMenu_auto ipsHide">
                            <form accept-charset="utf-8" method="post" action="//kingdom-leaks.com/index.php?/login/"
                                data-controller="core.global.core.login">
                                <input type="hidden" name="csrfKey" value="538d6993577ac1a4f5e0fc418dc51e57" />
                                <input type="hidden" name="ref"
                                    value="aHR0cHM6Ly9raW5nZG9tLWxlYWtzLmNvbS9pbmRleC5waHA/L2ZvcnVtcy90b3BpYy8zNzg5OS1wYXJhZGlzZS1sb3N0LW9ic2lkaWFuLWxpbWl0ZWQtZWRpdGlvbi0yMDIwLw==" />
                                <div data-role="loginForm">
                                    <div class="ipsPad ipsForm ipsForm_vertical">
                                        <h4 class="ipsType_sectionHead">Sign In</h4>
                                        <br /><br />
                                        <ul class="ipsList_reset">
                                            <li class="ipsFieldRow ipsFieldRow_noLabel ipsFieldRow_fullWidth">
                                                <input type="text" placeholder="Display Name" name="auth" />
                                            </li>
                                            <li class="ipsFieldRow ipsFieldRow_noLabel ipsFieldRow_fullWidth">
                                                <input type="password" placeholder="Password" name="password" />
                                            </li>
                                            <li class="ipsFieldRow ipsFieldRow_checkbox ipsClearfix">
                                                <span class="ipsCustomInput">
                                                    <input type="checkbox" name="remember_me" id="remember_me_checkbox"
                                                        value="1" checked="" aria-checked="true" />
                                                    <span></span>
                                                </span>
                                                <div class="ipsFieldRow_content">
                                                    <label class="ipsFieldRow_label" for="remember_me_checkbox">Remember
                                                        me</label>
                                                    <span class="ipsFieldRow_desc">Not recommended on shared
                                                        computers</span>
                                                </div>
                                            </li>
                                            <li class="ipsFieldRow ipsFieldRow_fullWidth">
                                                <br />
                                                <button type="submit" name="_processLogin" value="usernamepassword"
                                                    class="ipsButton ipsButton_primary ipsButton_small"
                                                    id="elSignIn_submit">Sign In</button>
                                                <br />
                                                <p class="ipsType_right ipsType_small">
                                                    <a href="https://kingdom-leaks.com/index.php?/lostpassword/"
                                                        data-ipsdialog="" data-ipsdialog-title="Forgot your password?">
                                                        Forgot your password?</a>
                                                </p>
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </li>
                    <li>
                        <a href="https://kingdom-leaks.com/index.php?/register/" id="elRegisterButton"
                            class="ipsButton ipsButton_normal ipsButton_primary">
                            Sign Up
                        </a>
                    </li>
                </ul>
            </div>
        </nav>
    </div>
    <main role="main" id="ipsLayout_body" class="ipsLayout_container v-nav-wrap ">
        <!--  Body Slider -->
        <div id="ipsLayout_contentArea">
            <div id="ipsLayout_contentWrapper">
                <nav class="ipsBreadcrumb ipsBreadcrumb_1 ipsFaded_withHover">
                    <a href="#" id="elRSS" class="ipsPos_right ipsType_large" title="Available RSS feeds"
                        data-ipstooltip="" data-ipsmenu="" data-ipsmenu-above=""><i class="fa fa-rss-square"></i></a>
                    <ul id="elRSS_menu" class="ipsMenu ipsMenu_auto ipsHide">
                        <li class="ipsMenu_item"><a title="The Kingdom Leaks Everything Feed"
                                href="https://kingdom-leaks.com/index.php?/rss/3-the-kingdom-leaks-everything-feed.xml/">The
                                Kingdom Leaks Everything Feed</a></li>
                        <li class="ipsMenu_item"><a title="New Downloads Only Feed"
                                href="https://kingdom-leaks.com/index.php?/rss/1-new-downloads-only-feed.xml/">New
                                Downloads Only Feed</a></li>
                        <li class="ipsMenu_item"><a title="New Downloads Only (without Singles) Feed"
                                href="https://kingdom-leaks.com/index.php?/rss/5-new-downloads-only-without-singles-feed.xml/">New
                                Downloads Only (without Singles) Feed</a></li>
                        <li class="ipsMenu_item"><a title="Newsroom Feed"
                                href="https://kingdom-leaks.com/index.php?/rss/2-newsroom-feed.xml/">Newsroom Feed</a>
                        </li>
                    </ul>
                    <ul class="ipsList_inline ipsPos_right">
                        <li>
                            <a class="ipsType_light" href="https://status.kingdom-leaks.com/" target="_blank"><i
                                    class="fa fa-server" aria-hidden="true"></i> <span>KLStatus</span></a>
                        </li>
                    </ul>
                    <ul class="ipsList_inline ipsPos_right">
                        <li>
                            <a class="ipsType_light" href="/release-notes/"><i class="fa fa-sticky-note"
                                    aria-hidden="true"></i> <span>Release Notes </span></a>
                        </li>
                    </ul>
                    <ul class="ipsList_inline ipsPos_right">
                        <li>
                            <a class="ipsType_light" href="/new/"><i class="fa fa-newspaper-o" aria-hidden="true"></i>
                                <span>Latest Downloads </span></a>
                        </li>
                    </ul>
                    <ul data-role="breadcrumbList">
                        <li>
                            <a title="Home" href="https://kingdom-leaks.com/">
                                <span><i class="fa fa-home"></i> Home <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            <a href="https://kingdom-leaks.com/index.php?/forums/">
                                <span>Forums <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            <a href="https://kingdom-leaks.com/index.php?/forums/forum/108-leaks/">
                                <span>Leaks <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/">
                                <span>Currently Leaked <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            Paradise Lost - Obsidian (Limited Edition) (2020)
                        </li>
                    </ul>
                </nav>
                <!--INSERT COMMUNITY MESSAGE HERE-->
                <!---->
                <div id="communityMessage">

                </div>
                <div class="cWidgetContainer " data-role="widgetReceiver" data-orientation="horizontal"
                    data-widgetarea="header">
                    <ul class="ipsList_reset">
                        <li class="ipsWidget ipsWidget_horizontal ipsBox"
                            data-blockid="app_featuredcontent_fcontentWidget_neazdnkpy" data-blockconfig="true"
                            data-blocktitle="Slider widget" data-blockerrormessage="This block cannot be shown. This
could be because it needs configuring, is unable to show on this page, or will show after reloading this page."
                            data-controller="core.front.widgets.block">
                            <ul class="ipsList_reset">
                                <div id="Fcontent_1">
                                    <div class="ipsPageHeader ipsClearfix sliderTitle">
                                        <h1 class="ipsType_pageTitle">Trending Leaks</h1>
                                        <a id="singles_page"
                                            style="font-size: 20px;right: 6px;top: -5px;color: #bfbfbf;display: none;"
                                            href="/index.php?/singles/" class="ipsRichEmbed_openItem" data-ipstooltip=""
                                            _title="View all Latest Singles"><i
                                                class="fa fa-external-link-square"></i></a>
                                        <a id="trending_page"
                                            style="font-size: 20px;right: 6px;top: -5px;color: #bfbfbf;display: none;"
                                            href="/index.php?/trending/" class="ipsRichEmbed_openItem"
                                            data-ipstooltip="" _title="View all Trending Leaks"><i
                                                class="fa fa-external-link-square">
                                                < /i></a>
                                        <a id="newsroom_forum"
                                            style="font-size: 20px;right: 6px;top: -5px;color: #bfbfbf;display: none;"
                                            href="/index.php?/forums/forum/45-the-newsroom/"
                                            class="ipsRichEmbed_openItem" data-ipstooltip=""
                                            _title="Visit The Newsroom"><i class="fa fa-externa
l-link-square"></i></a>
                                        <a id="discography_forum"
                                            style="font-size: 20px;right: 6px;top: -5px;color: #bfbfbf;display: none;"
                                            href="/index.php?/forums/forum/63-discographies/"
                                            class="ipsRichEmbed_openItem" data-ipstooltip=""
                                            _title="View all Latest Discographies">
                                            <i class="fa fa-external-link-square"></i></a>
                                        <a id="currently_leaked"
                                            style="font-size: 20px;right: 6px;top: -5px;color: #bfbfbf;display: none;"
                                            href="/index.php?/forums/forum/109-currently-leaked/"
                                            class="ipsRichEmbed_openItem" data-ipstooltip=""
                                            _title="View Currently Leaked forum"><i
                                                class="fa fa-external-link-square"></i></a>
                                        <a id="curated_playlists"
                                            style="font-size: 20px;right: 6px;top: -5px;color: #bfbfbf;display: none;"
                                            href="/index.php?/forums/forum/110-curated-playlists/"
                                            class="ipsRichEmbed_openItem" data-ipstooltip=""
                                            _title="View all Curated Playlists"><i cla
                                                ss="fa fa-external-link-square"></i></a>
                                    </div>
                                    <div id="sliderWrapper_1" class="sliderWrapper  nolinks norewrite">
                                        <ul class="slider_1">
                                            <li class="fcitem" id="fcc_37899">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/"
                                                    title="Paradise Lost - Obsidian (Limited Edition) (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is5-ssl.mzstatic.com/image/thumb/Music114/v4/f4/f8/67/f4f8672e-b25b-5185-7b8d-27f54fb
2af7c/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Paradise Lost - Obsidian (Limited Edition) (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Paradise Lost - Obsidian (Limited Edition) (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37802">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37802-killswitch-engage-atonement-ii-b-sides-for-charity-ep-2020/"
                                                    title="Killswitch Engage - Atonement II B-Sides for Charity [EP] (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i111.fastpic.ru/big/2020/0501/3b/1bbeb2713c42f2c551321ecbc20eec3b.jpg&amp;w=150&amp;h
=150&amp;pwd=quick@thumb!&amp;q=90" style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Killswitch Engage - Atonement II B-Sides for Charity [EP] (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Killswitch Engage - Atonement II B-Sides for Charity [EP]
                                                            (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37796">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37796-drake-dark-lane-demo-tapes-2020/"
                                                    title="Drake - Dark Lane Demo Tapes (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/FpTCyq.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="Drake - Dark Lane Demo Tapes (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Drake - Dark Lane Demo Tapes (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37690">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37690-our-last-night-let-light-overcome-the-darkness-2020/"
                                                    title="Our Last Night - Let Light Overcome The Darkness (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/3MkIQP.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="Our Last Night - Let Light Overcome The Darkness (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Our Last Night - Let Light Overcome The Darkness
                                                            (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37658">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37658-umbra-vitae-shadow-of-life-2020/"
                                                    title="Umbra Vitae - Shadow of Life (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/ahjn9s.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="Umbra Vitae - Shadow of Life (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Umbra Vitae - Shadow of Life (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37650">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37650-havok-v-2020/"
                                                    title="Havok - V (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/KxXQjn.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="Havok - V (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Havok - V (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37647">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37647-an-autumn-for-crippled-children-all-fell-silent-everything-went-quiet-2020/"
                                                    title="An Autumn for Crippled Children - All Fell Silent, Everyth
ing Went Quiet (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/64flEZ.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="An Autumn for Crippled Children - All Fell Silent, Everything Went Quiet (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>An Autumn for Crippled Children - All Fell Silent,
                                                            Everything Went Quiet (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37644">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37644-vader-solitude-in-madness-2020/"
                                                    title="Vader - Solitude in Madness (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/nvUXt6.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="Vader - Solitude in Madness (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Vader - Solitude in Madness (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37553">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37553-car-seat-headrest-making-a-door-less-open-2020/"
                                                    title="Car Seat Headrest - Making a Door Less Open (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is2-ssl.mzstatic.com/image/thumb/Music123/v4/b5/73/74/b5737457-c0af-1935-2187-4b43fb6
131c0/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Car Seat Headrest - Making a Door Less Open (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Car Seat Headrest - Making a Door Less Open (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37524">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37524-parkway-drive-viva-the-underdogs-the-film-2020/"
                                                    title="Parkway Drive - Viva the Underdogs [The Film] (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.vgy.me/9f9GG8.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height:
100%; max-height:150px; width:100%;" alt="Parkway Drive - Viva the Underdogs [The Film] (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Parkway Drive - Viva the Underdogs [The Film]
                                                            (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37506">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37506-boston-manor-glue-2020/"
                                                    title="Boston Manor - GLUE (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is4-ssl.mzstatic.com/image/thumb/Music123/v4/e9/4e/62/e94e62c4-2622-d754-9094-73113b2
29ed4/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Boston Manor - GLUE (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Boston Manor - GLUE (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37470">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37470-hayley-williams-petals-for-armor-2020/"
                                                    title="Hayley Williams - Petals For Armor (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://kingdom-leaks.com/uploads/monthly_2020_04/cover.jpg.d936098d0aaaedc4030ed15897e18a58.
jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90" style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Hayley Williams - Petals For Armor (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Hayley Williams - Petals For Armor (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37457">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37457-new-core-post-hardcoremetalcore-41720-42420/"
                                                    title="New Core: Post-Hardcore/Metalcore (4/17/20-4/24/20)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://kingdom-leaks.com/uploads/monthly_2020_04/041720.jpg.594e8b7304cd3c886e8e010c8b351866
.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90" style="height: 100%; max-height:150px; width:100%;"
                                                        alt="New Core: Post-Hardcore/Metalcore (4/17/20-4/24/20)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>New Core: Post-Hardcore/Metalcore (4/17/20-4/24/20)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37211">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37211-awolnation-angel-miners-the-lightning-riders-2020/"
                                                    title="AWOLNATION - Angel Miners &amp; the Lightning Riders (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is4-ssl.mzstatic.com/image/thumb/Music123/v4/c2/7b/c1/c27bc1cd-2fb0-9df0-56b7-3099e66
edb4e/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="AWOLNATION - Angel Miners &amp; the Lightning Riders (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>AWOLNATION - Angel Miners &amp; the Lightning Riders
                                                            (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37202">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37202-warbringer-weapons-of-tomorrow-2020/"
                                                    title="Warbringer - Weapons of Tomorrow (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is3-ssl.mzstatic.com/image/thumb/Music114/v4/db/d6/ab/dbd6ab8e-0a6c-1bbc-0d44-8dfff65
52eea/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Warbringer - Weapons of Tomorrow (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Warbringer - Weapons of Tomorrow (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37161">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37161-picturesque-do-you-feel-ok-2020/"
                                                    title="Picturesque - Do You Feel O.K (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is2-ssl.mzstatic.com/image/thumb/Music114/v4/f0/92/0f/f0920f42-79cf-43fb-080a-7ff3c30
1923c/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Picturesque - Do You Feel O.K (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Picturesque - Do You Feel O.K (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37158">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37158-the-used-heartwork-2020/"
                                                    title="The Used - Heartwork (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is1-ssl.mzstatic.com/image/thumb/Music113/v4/f3/98/dc/f398dc52-ef31-fe5c-8c9b-93edb16
089b0/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="The Used - Heartwork (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>The Used - Heartwork (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37123">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37123-dance-gavin-dance-afterburner-2020/"
                                                    title="Dance Gavin Dance - Afterburner (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is5-ssl.mzstatic.com/image/thumb/Music113/v4/4a/ef/5e/4aef5ecd-1f72-2f33-0075-063bad8
91f91/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Dance Gavin Dance - Afterburner (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Dance Gavin Dance - Afterburner (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37122">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37122-trivium-what-the-dead-men-say-japanese-edition-2020/"
                                                    title="Trivium - What the Dead Men Say (Japanese Edition) (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is5-ssl.mzstatic.com/image/thumb/Music123/v4/2d/e5/2b/2de52bc7-d27e-8552-ea20-7921af6
723c3/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Trivium - What the Dead Men Say (Japanese Edition) (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Trivium - What the Dead Men Say (Japanese Edition)
                                                            (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37114">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37114-i-am-abomination-passion-of-the-heist-ii-2020/"
                                                    title="I Am Abomination - Passion Of The Heist II (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is4-ssl.mzstatic.com/image/thumb/Music124/v4/c1/0e/74/c10e7453-0b81-787a-1ce2-00e7365
6a858/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="I Am Abomination - Passion Of The Heist II (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>I Am Abomination - Passion Of The Heist II (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37050">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37050-new-core-post-hardcoremetalcore-41020-41720/"
                                                    title="New Core: Post-Hardcore/Metalcore (4/10/20-4/17/20)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/mwQVBEW.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="New Core: Post-Hardcore/Metalcore (4/10/20-4/17/20)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>New Core: Post-Hardcore/Metalcore (4/10/20-4/17/20)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37037">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37037-mick-gordon-doom-eternal-original-game-soundtrack-2020/"
                                                    title="Mick Gordon - DOOM Eternal (Original Game Soundtrack) (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://kingdom-leaks.com/uploads/monthly_2020_04/cover_1_600x600.jpg.cbbf1ef6cd0ddfc22b2b1aa
9486e155c.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90" style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Mick Gordon - DOOM Eternal (Original Game Soundtrack) (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Mick Gordon - DOOM Eternal (Original Game Soundtrack)
                                                            (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37024">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37024-moby-all-visible-objects-2020/"
                                                    title="Moby - All Visible Objects (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is5-ssl.mzstatic.com/image/thumb/Music113/v4/9a/b5/d6/9ab5d664-1b17-8c9a-a155-068c820
b6ac5/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Moby - All Visible Objects (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Moby - All Visible Objects (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_37012">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/37012-katatonia-city-burials-2020/"
                                                    title="Katatonia - City Burials (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is1-ssl.mzstatic.com/image/thumb/Music123/v4/48/da/dd/48daddb8-38a7-74e2-b271-18b6476
ebd44/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Katatonia - City Burials (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Katatonia - City Burials (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36874">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36874-fiona-apple-fetch-the-bolt-cutters-2020/"
                                                    title="Fiona Apple - Fetch The Bolt Cutters (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/n8OYztM.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Fiona Apple - Fetch The Bolt Cutters (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Fiona Apple - Fetch The Bolt Cutters (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36866">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36866-the-acacia-strain-c-singles-2020/"
                                                    title="The Acacia Strain - C (Singles) (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/6t7isUg.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="The Acacia Strain - C (Singles) (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>The Acacia Strain - C (Singles) (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36842">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36842-pulses-speak-it-into-existence-2020/"
                                                    title="Pulses. - Speak It Into Existence (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/QUaQfss.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Pulses. - Speak It Into Existence (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Pulses. - Speak It Into Existence (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36825">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36825-enter-shikari-nothing-is-true-everything-is-possible-instrumentals-2020/"
                                                    title="Enter Shikari - Nothing is True &amp; Everything is Possible
(+Instrumentals) (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/BXu0M39.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Enter Shikari - Nothing is True &amp; Everything is Possible (+Instrumentals) (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Enter Shikari - Nothing is True &amp; Everything is
                                                            Possible (+Instrumentals) (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36718">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36718-a-day-to-remember-mindreader-single-2020/"
                                                    title="A Day to Remember - Mindreader [Single] (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i111.fastpic.ru/big/2020/0414/ee/0f5bf5d1799052fd8d574d3efce7e4ee.jpg&amp;w=150&amp;h
=150&amp;pwd=quick@thumb!&amp;q=90" style="height: 100%; max-height:150px; width:100%;"
                                                        alt="A Day to Remember - Mindreader [Single] (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>A Day to Remember - Mindreader [Single] (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36667">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36667-letlive-10-years-of-fake-history-demos-2020/"
                                                    title="letlive. - 10 Years Of Fake History (Demos) (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://f4.bcbits.com/img/a3246297941_10.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=9
0" style="height: 100%; max-height:150px; width:100%;" alt="letlive. - 10 Years Of Fake History (Demos) (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>letlive. - 10 Years Of Fake History (Demos) (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36553">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36553-emery-white-line-fever-2020/"
                                                    title="Emery - White Line Fever (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/eEgMqBo.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Emery - White Line Fever (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Emery - White Line Fever (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36546">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36546-grey-daze-amends-2020/"
                                                    title="Grey Daze - Amends (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is5-ssl.mzstatic.com/image/thumb/Music123/v4/d2/18/e2/d218e22c-79f6-e922-170f-87319ed
85aa6/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Grey Daze - Amends (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Grey Daze - Amends (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36487">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36487-kill-the-lights-plagues-ep-2020/"
                                                    title="Kill the Lights - Plagues [EP] (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/DluBxoa.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Kill the Lights - Plagues [EP] (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Kill the Lights - Plagues [EP] (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36453">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36453-rotting-out-ronin-2020/"
                                                    title="Rotting Out - Ronin (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/oALbM5q.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Rotting Out - Ronin (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Rotting Out - Ronin (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36452">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36452-drain-california-cursed-2020/"
                                                    title="DRAIN - California Cursed (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/lry95SN.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="DRAIN - California Cursed (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>DRAIN - California Cursed (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36449">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36449-the-strokes-the-new-abnormal-2020/"
                                                    title="The Strokes - The New Abnormal (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/3rYHEm8.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="The Strokes - The New Abnormal (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>The Strokes - The New Abnormal (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36421">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36421-sparta-trust-the-river-2020/"
                                                    title="Sparta - Trust the River (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/DcZIFc1.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Sparta - Trust the River (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Sparta - Trust the River (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36410">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36410-dream-on-dreamer-what-if-i-told-you-it-doesnt-get-better-2020/"
                                                    title="Dream On Dreamer - What If I Told You It Doesn&#39;t Get Better (2020)">
                                                    <img class="lazyload" src="/img/filler.png"
                                                        data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://i.imgur.com/uNLoEeG.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="hei
ght: 100%; max-height:150px; width:100%;" alt="Dream On Dreamer - What If I Told You It Doesn&#39;t Get Better (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Dream On Dreamer - What If I Told You It Doesn&#39;t Get
                                                            Better (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36369">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36369-local-h-lifers-2020/"
                                                    title="Local H - Lifers (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://is3-ssl.mzstatic.com/image/thumb/Music113/v4/1a/35/29/1a35292f-eb3d-b6c9-56fb-176acfb
4a977/source/600x600bb.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90"
                                                        style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Local H - Lifers (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Local H - Lifers (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                            <li class="fcitem" id="fcc_36216">
                                                <a href="https://kingdom-leaks.com/index.php?/forums/topic/36216-breakdowns-at-tiffanys-eternal-lords-2020/"
                                                    title="Breakdowns At Tiffany&#39;s - Eternal Lords (2020)">
                                                    <img class="lazyload" src="/img/filler.png" data-src="https://kingdom-leaks.com/thumb/webroot/img.php?src=https://metalnoise.net/wp-content/uploads/2019/10/69151952_2356214144434237_26892214880908083
2_o.jpg&amp;w=150&amp;h=150&amp;pwd=quick@thumb!&amp;q=90" style="height: 100%; max-height:150px; width:100%;"
                                                        alt="Breakdowns At Tiffany&#39;s - Eternal Lords (2020)" />
                                                    <div class="bx-caption" style="display: block;">
                                                        <span>Breakdowns At Tiffany&#39;s - Eternal Lords (2020)</span>
                                                    </div>
                                                </a>
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </ul>
                        </li>
                        <li class="ipsWidget ipsWidget_horizontal ipsBox ipsWidgetHide ipsHide"
                            data-blockid="app_portal_portalBlock_cd2g0ltfq" data-blockconfig="true"
                            data-blocktitle="Portal Blocks"
                            data-blockerrormessage="This block cannot be sh
own. This could be because it needs configuring, is unable to show on this page, or will show after reloading this page."
                            data-controller="core.front.widgets.block"></li>
                        <li class="ipsWidget ipsWidget_horizontal ipsBox ipsWidgetHide ipsHide"
                            data-blockid="app_bimchatbox_bimchatbox_h87rdnpit" data-blocktitle="Chatbox"
                            data-blockerrormessage="This block cannot be shown. This could be because
it is unable to show on this specific page, or will show after reloading this page."
                            data-controller="core.front.widgets.block"></li>
                        <li class="ipsWidget ipsWidget_horizontal ipsBox"
                            data-blockid="app_portal_portalBlock_z9be6yke9" data-blockconfig="true"
                            data-blocktitle="Portal Blocks" data-blockerrormessage="This block cannot be shown. This could be bec
ause it needs configuring, is unable to show on this page, or will show after reloading this page."
                            data-controller="core.front.widgets.block"><a id="togglePlaylists" href="#"
                                style="text-align:center;margin:auto;display:block;padding: 6px;color:#bfbfbf"><i
                                    class="fa fa-caret-down" aria-hidden="true"></i> Curated Playlists by Kingdom Leaks
                                <i class="fa fa-caret-down" aria-hidden="true"></i></a>
                            <div id="playlistContainer">
                                <div id="curatedPlaylists">
                                    <div class="playlist"><img data-ipstooltip=""
                                            title="Playlist of new post-hardcore/metalcore releases posted on Kingdom Leaks. Updates every weekend."
                                            src="https://kingdom-leaks.com/img/phc-mc-new.jpg"
                                            class="playlist-image" /><br />
                                        <span class="ipsBadge ipsBadge_style6" style="margin-bottom: -12px;">LISTEN //
                                            LINKS</span>
                                        <p style="text-align: center;">
                                            <a data-ipstooltip="" title="Listen on Spotify"
                                                href="https://open.spotify.com/playlist/4KJOjGo1WwjEXIwLPUPDL1"
                                                target="_blank">
                                                <img alt="Spotify" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/spotify.png"
                                                    style="width: 25px; height: auto;" /></a>
                                            <a data-ipstooltip="" title="Listen on Apple Music"
                                                href="https://music.apple.com/us/playlist/new-core-post-hardcore-metalcore/pl.u-zlGXgCdkRRWk"
                                                target="_blank">
                                                <img alt="Apple Music" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/am.png"
                                                    style="width: 25px; height: auto;" /></a>
                                            <a data-ipstooltip="" title="Listen on TIDAL"
                                                href="https://listen.tidal.com/playlist/7d46865a-1776-4d0e-995b-5801f7467a7f"
                                                target="_blank">
                                                <img alt="tidal.png" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/tidal.png"
                                                    style="width: 31px; height: auto;" /></a>
                                            <a data-ipstooltip="" title="Listen on YouTube/YouTube Music"
                                                href="https://www.youtube.com/playlist?list=PLm__V4sSbIiMJw0nhq117aaFB6HcUtCmb"
                                                target="_blank">
                                                <img alt="ym.png" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/youtube.png"
                                                    style="width: 27px; height: auto;" /></a>
                                        </p>
                                    </div>
                                    <div class="playlist"><img data-ipstooltip=""
                                            title="Archive of all previously featured post-hardcore/metalcore tracks posted on Kingdom Leaks."
                                            src="https://kingdom-leaks.com/img/phc-mc-archive.jpg"
                                            class="playlist-image" /><br />
                                        <span class="ipsBadge ipsBadge_style6" style="margin-bottom: -12px;">LISTEN //
                                            LINKS</span>
                                        <p style="text-align: center;">
                                            <a data-ipstooltip="" title="Listen on Spotify"
                                                href="https://open.spotify.com/playlist/0hisjoUYvFLaUcDANw5se3"
                                                target="_blank">
                                                <img alt="Spotify" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/spotify.png"
                                                    style="width: 25px; height: auto;" /></a>
                                            <a data-ipstooltip="" title="Listen on Apple Music"
                                                href="https://music.apple.com/us/playlist/new-core-post-hardcore-metalcore-2020-archive/pl.u-5Z4x6FxgaaJg"
                                                target="_blank">
                                                <img alt="Apple Music" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/am.png"
                                                    style="width: 25px; height: auto;" /></a>
                                            <a data-ipstooltip="" title="Listen on TIDAL"
                                                href="https://listen.tidal.com/playlist/38ce02d6-bb2c-48ba-ae2d-ffe7ee9ea296"
                                                target="_blank">
                                                <img alt="tidal.png" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/tidal.png"
                                                    style="width: 31px; height: auto;" /></a>
                                            <a data-ipstooltip="" title="Listen on YouTube/YouTube Music"
                                                href="https://www.youtube.com/playlist?list=PLm__V4sSbIiMJw0nhq117aaFB6HcUtCmb"
                                                target="_blank">
                                                <img alt="ym.png" class="ipsImage"
                                                    src="https://kingdom-leaks.com/img/youtube.png"
                                                    style="width: 27px; height: auto;" /></a>
                                        </p>
                                    </div>
                                    <span
                                        style="width: 100%; display: inline-block; font-size: 0; line-height: 0;"></span>
                                </div>
                            </div>
                            <script>
                                window.addEventListener('DOMContentLoaded', (event) => {
                                    document.getElementById("togglePlaylists").addEventListener("click", function (e) {
                                        e.preventDefault();
                                        $("#playlistContainer").fadeToggle("fast", "linear", function () {
                                            if (document.getElementById("playlistContainer").style.display == "block") {
                                                document.getElementById("togglePlaylists").innerHTML = '<i class="fa fa-caret-up" aria-hidden="true"></i> Curated Playlists by Kingdom Leaks <i class="fa fa-caret-up" aria-hidden="true"></i>'
                                            } else {
                                                document.getElementById("togglePlaylists").innerHTML = '<i class="fa fa-caret-down" aria-hidden="true"></i> Curated Playlists by Kingdom Leaks <i class="fa fa-caret-down" aria-hidden="true"></i>'
                                            }
                                        });
                                    });
                                });
                            </script>
                        </li>
                    </ul>
                </div>
                <!-- Toggle buttons -->
                <!-- endif -->
                <div id="ipsLayout_mainArea" style="display:table-cell;">
                    <a id="elContent"></a>
                    <div class="ipsPageHeader ipsClearfix">
                        <div class="ipsPos_right ipsResponsive_noFloat ipsResponsive_hidePhone">
                            <div data-followapp="forums" data-followarea="topic" data-followid="37899"
                                data-controller="core.front.core.followButton">
                                <span
                                    class="ipsType_light ipsType_blendLinks ipsResponsive_hidePhone ipsResponsive_inline"><i
                                        class="fa fa-info-circle"></i> <a
                                        href="https://kingdom-leaks.com/index.php?/login/"
                                        title="Go to the sign in page">Sign in to follow this</a>  </span>
                                <div class="ipsFollow ipsPos_middle ipsButton ipsButton_link ipsButton_verySmall ipsButton_disabled"
                                    data-role="followButton">
                                    <span>Followers</span>
                                    <span class="ipsCommentCount">0</span>
                                </div>
                            </div>
                        </div>
                        <div class="ipsPos_right ipsResponsive_noFloat ipsResponsive_hidePhone">
                        </div>
                        <div class="ipsPhotoPanel ipsPhotoPanel_small ipsPhotoPanel_notPhone ipsClearfix">
                            <a href="https://kingdom-leaks.com/index.php?/profile/57459-summers/" data-ipshover=""
                                data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/57459-summers/&amp;do=hovercard"
                                class="ipsUserPhoto ipsUserPhoto_small" title="Go to Summers&#39;s p
rofile">
                                <img src="https://kingdom-leaks.com/uploads/monthly_2020_04/imageproxy.thumb.jpg.20df163213bc6f362a772b5aad24d2d5.jpg"
                                    alt="Summers" itemprop="image" />
                            </a>
                            <div>
                                <h1 class="ipsType_pageTitle ipsContained_container">
                                    <span><span class="ipsBadge ipsBadge_icon ipsBadge_positive" data-ipstooltip=""
                                            title="Featured"><i class="fa fa-star"></i></span></span>
                                    <div id="releaseTitle" class="ipsType_break ipsContained">
                                        Paradise Lost - Obsidian (Limited Edition) (2020)
                                    </div>
                                </h1>
                                <p class="ipsType_reset ipsType_blendLinks ">
                                    <span class="ipsType_normal">Started by
                                        <a href="https://kingdom-leaks.com/index.php?/profile/57459-summers/"
                                            data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/57459-summers/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Ffor
ums%252Ftopic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Summers&#39;s profile"
                                            class="ipsType_break"><span style="color:#f39c12"><i
                                                    class="fa fa-shield"></i> Summers</span></a>
                                    </span>, <span class="ipsType_light ipsType_noBreak"><time
                                            datetime="2020-05-04T17:00:56Z" title="05/04/2020 05:00  PM"
                                            data-short="35 min">35 minutes ago</time></span><br />
                                </p>
                            </div>
                        </div>
                    </div>
                    <div class="ipsClearfix">
                        <ul
                            class="ipsToolList ipsToolList_horizontal ipsClearfix ipsSpacer_both ipsResponsive_hidePhone">

                        </ul>
                    </div>
                    <div data-controller="core.front.core.commentFeed,forums.front.topic.view, core.front.core.ignoredComments"
                        data-autopoll=""
                        data-baseurl="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/"
                        data-lastpage="" data- feedid="topic-37899" class="cTopic ipsClear ipsSpacer_top">
                        <h2 class="ipsType_sectionTitle ipsType_reset ipsType_medium" data-role="comment_count"
                            data-commentcountstring="js_num_topic_posts">1 post in this topic</h2>
                        <div data-role="commentFeed" data-controller="core.front.core.moderation"
                            class="ipsAreaBackground_light">
                            <form
                                action="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/&amp;csrfKey=538d6993577ac1a4f5e0fc418dc51e57&amp;do=multimodComment"
                                method="post" data-ipspageaction="" data-role="moderationTools">
                                <a id="comment-229357"></a>
                                <article id="elComment_229357"
                                    class="cPost ipsBox  ipsComment  ipsComment_parent ipsClearfix ipsClear ipsColumns ipsColumns_noSpacing ipsColumns_collapsePhone  ">
                                    <div class="cAuthorPane cAuthorPane_mobile ipsResponsive_showPhone ipsResponsive_block"
                                        style="display:none;">
                                        <h3
                                            class="ipsType_sectionHead cAuthorPane_author ipsResponsive_showPhone ipsResponsive_inlineBlock ipsType_break ipsType_blendLinks ipsTruncate ipsTruncate_line">
                                            <a href="https://kingdom-leaks.com/index.php?/profile/57459-summers/"
                                                data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/57459-summers/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Ffor
ums%252Ftopic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Summers&#39;s profile"
                                                class="ipsType_break"><span style="color:#f39c12"><i
                                                        class="fa fa-shield"></i> Summers</span></a>
                                            <span class="ipsResponsive_showPhone ipsResponsive_inline">  
                                                <span title="Member&#39;s total reputation" data-ipstooltip=""
                                                    class="ipsRepBadge ipsRepBadge_positive">
                                                    <i class="fa fa-plus-circle"></i> 8,137
                                                </span>
                                            </span>
                                        </h3>
                                        <div class="lastfmMobile lazyload"></div>
                                        <div class="cAuthorPane_photo">
                                            <a href="https://kingdom-leaks.com/index.php?/profile/57459-summers/"
                                                data-ipshover=""
                                                data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/57459-summers/&amp;do=hovercard"
                                                class="ipsUserPhoto ipsUserPhoto_large" title="Go to Summers&#39;s p
rofile">
                                                <img src="https://kingdom-leaks.com/uploads/monthly_2020_04/imageproxy.thumb.jpg.20df163213bc6f362a772b5aad24d2d5.jpg"
                                                    alt="Summers" itemprop="image" />
                                            </a>
                                        </div>
                                    </div>
                                    <aside
                                        class="ipsComment_author cAuthorPane ipsColumn ipsColumn_medium euip_PanelWidth">
                                        <div class="euip_mobile ipsResponsive_hidePhone">
                                            <h3 class="ipsType_sectionHead cAuthorPane_author ipsType_blendLinks ipsType_break euip_UserNameFont"
                                                itemprop="creator" itemscope="" itemtype="http://schema.org/Person">
                                                <i class="fa fa-circle euipOnlineStatus_offline" data-ipstooltip=""
                                                    title="Offline"></i>
                                                <strong itemprop="name">
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/57459-summers/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/57459-summers/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Ffor
ums%252Ftopic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Summers&#39;s profile"
                                                        class="ipsType_break"><span style="color:#f39c12"><i
                                                                class="fa fa-shield"></i> Summers</span></a></strong>
                                            </h3>
                                        </div>
                                        <ul class="cAuthorPane_info ipsList_reset ">
                                            <li class="euip_Photo ipsResponsive_hidePhone">
                                                <span class="ipsUserPhoto euip_AvatarSize ">
                                                    <img class="lazyload userOffline"
                                                        data-src="https://kingdom-leaks.com/uploads/monthly_2020_04/imageproxy.thumb.jpg.20df163213bc6f362a772b5aad24d2d5.jpg"
                                                        src="/img/filler.png" alt="Summers" itemprop="image" />
                                                </span>
                                            </li>
                                            <li class="ipsResponsive_hidePhone ipsType_break">
                                                Elite
                                            </li>
                                            <li class="ipsResponsive_hidePhone">
                                                <img src="https://kingdom-leaks.com/uploads/monthly_2017_11/Provi.png.fb351fdfd2b5e0282348bfe8d1584caa.png"
                                                    alt="" class="cAuthorGroupIcon" />
                                            </li>
                                            <div class="euip_InfoPanel">
                                                <hr class="euip_Hr ipsResponsive_hidePhone" />
                                                <li class="euip_Border ipsResponsive_hidePhone">
                                                    <span class="euip_Title"><i class="fa fa-keyboard-o"
                                                            aria-hidden="true"></i> Post Count:</span>
                                                    <span class="euip_Content">4,658</span>
                                                </li>
                                                <br class="ipsResponsive_hidePhone" />
                                                <hr class="euip_Hr ipsResponsive_hidePhone" />
                                                <li class="euip_Border ipsResponsive_hidePhone">
                                                    <span class="euip_Title"><i class="fa fa-thumbs-up"
                                                            aria-hidden="true"></i> Reputation:</span>
                                                    <span class="euip_Content">
                                                        <span title="Member&#39;s total reputation" data-ipstooltip=""
                                                            class="ipsRepBadge ipsRepBadge_positive">
                                                            <i class="fa fa-plus-circle"></i> 8,137
                                                        </span>
                                                    </span>
                                                </li>
                                                <br class="ipsResponsive_hidePhone" />
                                                <hr class="euip_Hr" />
                                                <li id="" class="ipsResponsive_hidePhone ipsType_break ">
                                                </li>
                                                <li class="ipsType_break euip_Border ipsResponsive_hidePhone">
                                                    <span class="euip_Title"><i class="fa fa-music"
                                                            aria-hidden="true"></i> Favorite Genre:</span>
                                                    <span class="euip_Content">Metalcore</span>
                                                </li>
                                                <br class"ipsresponsive_hidephone"="" />
                                                <hr class="euip_Hr ipsResponsive_hidePhone" />
                                                <li id="" class="ipsResponsive_hidePhone ipsType_break ">
                                                </li>
                                                <li class="ipsType_break euip_Border ipsResponsive_hidePhone">
                                                    <span class="euip_Title"><i class="fa fa-play"
                                                            aria-hidden="true"></i> Favorite Artist:</span>
                                                    <span class="euip_Content" data-ipstooltip=""
                                                        _title="TGI/Bury Tomorrow">TGI/Bury Tomorrow</span>
                                                </li>
                                                <br class="ipsResponsive_hidePhone" />
                                                <hr class="euip_Hr ipsResponsive_hidePhone" />
                                                <li id="" class="ipsResponsive_hidePhone ipsType_break ">
                                                </li>
                                                <li class="ipsType_break euip_Border ipsResponsive_hidePhone">
                                                    <span class="euip_Title"><i class="fa fa-archive"
                                                            aria-hidden="true"></i> Preferred Audio Format:</span>
                                                    <span class="euip_Content">MP3/320</span>
                                                </li>
                                                <br class="ipsResponsive_hidePhone" />
                                                <hr class="euip_Hr ipsResponsive_hidePhone" />
                                                <li class="ipsResponsive_hidePhone"><span class="ipsPip"></span><span
                                                        class="ipsPip"></span><span class="ipsPip"></span></li>
                                            </div>
                                        </ul>
                                    </aside>
                                    <div class="ipsColumn ipsColumn_fluid">
                                        <div id="comment-229357_wrap" data-controller="core.front.core.comment"
                                            data-commentapp="forums" data-commenttype="forums" data-commentid="229357"
                                            data-quotedata="{&#34;userid&#34;:57459,&#34;username&#34;:&#34;Summers&#34;,&#34;timestamp&#34;:1588611656,&#34;con
tentapp&#34;:&#34;forums&#34;,&#34;contenttype&#34;:&#34;forums&#34;,&#34;contentid&#34;:37899,&#34;contentclass&#34;:&#34;forums_Topic&#34;,&#34;contentcommentid&#34;:229357}"
                                            class="ipsComment_content ipsType_medium  ipsFaded_withHover">
                                            <div class="ipsComment_meta ipsType_light">
                                                <div
                                                    class="ipsPos_right ipsType_light ipsType_reset ipsFaded ipsFaded_more ipsType_blendLinks">
                                                    <ul class="ipsList_inline ipsComment_tools">
                                                        <li><a href="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/&amp;do=reportComment&amp;comment=229357"
                                                                data-ipsdialog="" data-ipsdialog-remotesubmit=""
                                                                data-ipsdialog-size="medium"
                                                                data-ipsdialog-flashmessage="Thanks for your report."
                                                                data-ipsdialog-title="Report post"
                                                                data-action="reportComment" title="Report this content"
                                                                class="ipsFaded ipsFaded_more"><span
                                                                    class="ipsResponsive_showPhone ipsResponsive_inline"><i
                                                                        class="fa fa-fl
ag"></i></span><span class="ipsResponsive_hidePhone ipsResponsive_inline ipsBadge ipsBadge_style4">REPORT
                                                                    POST</span></a></li>
                                                        <li><a class="ipsType_blendLinks"
                                                                href="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/"
                                                                data-ipstooltip="" title="Share this post"
                                                                data-ipsmenu="" data-ipsmenu-closeoncl ick="false"
                                                                id="elSharePost_229357" data-role="shareComment"><i
                                                                    class="fa fa-share-alt"></i></a></li>
                                                    </ul>
                                                </div>
                                                <div class="ipsType_reset">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/&amp;do=findComment&amp;comment=229357"
                                                        class="ipsType_blendLinks">Posted <time
                                                            datetime="2020-05-04T17:00:56Z" title="05/04/2020
05:00  PM" data-short="35 min">35 minutes ago</time></a>
                                                    <span class="ipsResponsive_hidePhone">
                                                    </span>
                                                </div>
                                            </div>
                                            <div class="cPost_contentWrap ipsPad">
                                                <div data-role="commentContent"
                                                    class="ipsType_normal ipsType_richText ipsContained"
                                                    data-controller="core.front.core.lightboxedImages">
                                                    <p style="text-align:center;">
                                                        <img alt="600x600bb.jpg" class="ipsImage" height="600" src="https://kingdom-leaks.com/applications/core/interface/imageproxy/imageproxy.php?img=https://is5-ssl.mzstatic.com/image/thumb/Music114/v4/f4/f8/67/f4f8672e-b25b-5185-7b8d-27f54fb2af7c/source/600x600bb.jpg&amp;key=a354a8baa5d145f970ae06fdea6bd929e5752c40a504258e8d750315ebfae55c" width="600"
                                                            data-imageproxy-source="https://is5-ssl.mzstatic.com/image/thumb/Music114/v4/f4/f8/67/f4f8672e-b25b-5185-7b8d-27f54fb2af7c/source/600x600bb.jpg" />
                                                    </p>
                                                    <p style="text-align:center;">
                                                        Leaked<br />
                                                        Release - May 15th, 2020
                                                    </p>
                                                    <p style="text-align:center;">
                                                        Genre - Gothic Metal, Doom Metal
                                                    </p>
                                                    <p style="text-align:center;">
                                                        Quality - MP3, 160 kbps CBR
                                                    </p>
                                                    <p>
                                                         
                                                    </p>
                                                    <p style="text-align:center;">
                                                        <strong>Tracklist:</strong>
                                                    </p>
                                                    <p style="text-align: center;">
                                                        <strong>1.</strong> Darker Thoughts <em>(05:46)</em><br />
                                                        <strong>2.</strong> Fall from Grace <em>(05:42)</em><br />
                                                        <strong>3.</strong> Ghosts <em>(04:35)</em><br />
                                                        <strong>4.</strong> The Devil Embraced <em>(06:08)</em><br />
                                                        <strong>5.</strong> Forsaken <em>(04:30)</em><br />
                                                        <strong>6.</strong> Serenity <em>(04:46)</em><br />
                                                        <strong>7.</strong> Ending Days <em>(04:36)</em><br />
                                                        <strong>8.</strong> Hope Dies Young <em>(04:02)</em><br />
                                                        <strong>9.</strong> Ravenghast <em>(05:30)</em><br />
                                                        <strong>10.</strong> Hear the Night <em>(05:34)</em><br />
                                                        <strong>11.</strong> Defiler <em>(04:45)</em>
                                                    </p>
                                                    <p>
                                                         
                                                    </p>
                                                    <p style="text-align:center;">
                                                        <strong><a
                                                                href="https://kingdom-leaks.com/index.php?/go/&amp;url=https://filecrypt.cc/Container/CC6F0BD926.html"
                                                                rel="external nofollow">Download</a></strong>
                                                    </p>
                                                    <p>
                                                         
                                                    </p>
                                                    <p style="text-align:center;">
                                                        <strong>Support!</strong>
                                                    </p>
                                                    <p style="text-align:center;">
                                                        <a href="https://kingdom-leaks.com/index.php?/go/&amp;url=https://www.facebook.com/paradiselostofficial/"
                                                            rel="external nofollow">Facebook</a> / <a href="https://kingdom-leaks.com/index.php?/go/&amp;url=https://music.apple.com/us/album/obsidian/1501148713
?uo=4" rel="external nofollow">iTunes</a>
                                                    </p>
                                                    <p style="text-align:center;">
                                                         
                                                    </p>
                                                    <div class="ipsEmbeddedVideo ipsEmbeddedVideo_limited"
                                                        contenteditable="false">
                                                        <div>
                                                            <a
                                                                href="https://kingdom-leaks.com/index.php?/go/&amp;url=https://www.youtube.com/watch?v=urugx_wSBy8">
                                                                <div class="lazyvideo" id="urugx_wSBy8" style="background-image: url(https://i.ytimg.com/vi/urugx_wSBy8/mqdefault.jpg); width: 100%; height: 100
%"></div>
                                                            </a>
                                                        </div>
                                                    </div>
                                                    <p style="text-align:center;">
                                                         
                                                    </p>
                                                </div>
                                                <div class="ipsItemControls">
                                                    <div data-controller="core.front.core.reaction"
                                                        class="ipsItemControls_right ipsClearfix ">
                                                        <div class="ipsReact ipsPos_right">
                                                            <div class="ipsReact_blurb " data-role="reactionBlurb">
                                                                <ul class="ipsReact_reactions">
                                                                    <li class="ipsReact_reactCount">
                                                                        <span data-ipstooltip="" title="Love">
                                                                            <span>
                                                                                <img src="https://kingdom-leaks.com/uploads/reactions/react_like.png"
                                                                                    alt="Love" />
                                                                            </span>
                                                                            <span>
                                                                                1
                                                                            </span>
                                                                        </span>
                                                                    </li>
                                                                </ul>
                                                                <div class="ipsReact_overview ipsType_blendLinks">

                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <ul class="ipsComment_controls ipsClearfix ipsItemControls_left"
                                                        data-role="commentControls">
                                                        <li class="ipsHide" data-role="commentLoading">
                                                            <span
                                                                class="ipsLoading ipsLoading_tiny ipsLoading_noAnim"></span>
                                                        </li>
                                                    </ul>
                                                </div>
                                            </div>
                                            <div class="ipsMenu ipsMenu_wide ipsHide cPostShareMenu"
                                                id="elSharePost_229357_menu">
                                                <div class="ipsPad">
                                                    <h4 class="ipsType_sectionHead">Share this post</h4>
                                                    <hr class="ipsHr" />
                                                    <h5 class="ipsType_normal ipsType_reset">Link to post</h5>
                                                    <input type="text"
                                                        value="https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/"
                                                        class="ipsField_fullWidth" />
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </article>
                                <!-- SIMILAR TOPICS -->
                                <div id="otherReleasesBy"
                                    class="cPost ipsBox  ipsComment  ipsComment_parent ipsClearfix ipsClear ipsColumns ipsColumns_noSpacing ipsColumns_collapsePhone">
                                    <h3 id="relatedTitle"
                                        style="text-align: center; padding: 11px 15px; font-size: 15px; font-weight: 400; color: #bfbfbf;"
                                        class="ipsWidget ipsWidget_horizontal ipsWidget_title ipsType_reset ipsBox">
                                        Other Releases by
                                        <a id="moreResults"
                                            style="font-size: 20px; right: 6px; top: 5px; color: #bfbfbf;"
                                            href="https://kingdom-leaks.com/index.php?/search/&amp;search_and_or=and&amp;search_in=titles&amp;sortby=newest&amp;q="
                                            class="ipsRichEmbed_openItem" target="_blank" data -ipstooltip=""
                                            _title="View all releases by this artist"><i
                                                class="fa fa-external-link-square"></i></a>
                                    </h3>
                                    <ul class="ipsDataList" style="background: #212121;">
                                        <li class="ipsDataItem ipsDataItem_responsivePhoto   ">
                                            <div class="bim_tthumb_wrap" style="padding: 0px !important;">
                                                <div class="bim_tthumb_wrap">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/17371-paradise-lost-believe-in-nothing-remixed-remastered-2018/"
                                                        data-ipstpopup="" data-ipstpopup-width="280"
                                                        data-ipstpopup-target="https://kingdom-leaks.com/index.php?/forums/top
ic/17371-paradise-lost-believe-in-nothing-remixed-remastered-2018/&amp;do=tthumbGetImage"
                                                        alt="Paradise Lost - Believe In Nothing [Remixed &amp; Remastered] (2018)">
                                                        <div data-full="https://kingdom-leaks.com/uploads/monthly_2018_06/topicThumb_17371.jpg.d4528eaf8f1989b2600fcef81a776e60.jpg"
                                                            id="tthumb_17371" class="tthumb_standard lazyload" data-bg="https://kingdom-leaks.com/uploads/monthly_2018
_06/topicThumb_17371.jpg.thumb.jpg.ba5178dccbad4f38fedcb2c32797818e.jpg" style="width: 60px; height: 60px;"></div>
                                                    </a>
                                                    <!---->
                                                </div>
                                            </div>
                                            <div class="ipsDataItem_main">
                                                <h4 class="ipsDataItem_title ipsContained_container">
                                                    <span class="ipsType_break ipsContained">
                                                        <a href="https://kingdom-leaks.com/index.php?/forums/topic/17371-paradise-lost-believe-in-nothing-remixed-remastered-2018/"
                                                            data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/forums/topic/17371-paradise
-lost-believe-in-nothing-remixed-remastered-2018/&amp;preview=1" data-ipshover-timeout="1.5">
                                                            Paradise Lost - Believe In Nothing [Remixed &amp;
                                                            Remastered] (2018)
                                                        </a>
                                                    </span>
                                                </h4>
                                                <p class="ipsType_reset ipsType_medium ipsType_light">
                                                    By
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/10013-d666n/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/10013-d666n/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Fforums%
252Ftopic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to D666N&#39;s profile"
                                                        class="ipsType_break"><span style="color:dimgrey"><i
                                                                class="fa fa-glass"></i> D666N</span></a>
                                                    <time datetime="2018-06-29T12:06:26Z" title="06/29/2018 12:06  PM"
                                                        data-short="1 yr">June 29, 2018</time>
                                                    in <a
                                                        href="https://kingdom-leaks.com/index.php?/forums/forum/11-metal/">Metal</a>
                                                </p>
                                                <ul class="ipsList_inline ipsClearfix ipsType_light">
                                                </ul>
                                            </div>
                                            <ul class="ipsDataItem_stats">
                                                <li>
                                                    <span class="ipsDataItem_stats_number">3</span>
                                                    <span class="ipsDataItem_stats_type"> replies</span>
                                                </li>
                                                <li>
                                                    <span class="ipsDataItem_stats_number">3,315</span>
                                                    <span class="ipsDataItem_stats_type"> views</span>
                                                </li>
                                            </ul>
                                            <ul class="ipsDataItem_lastPoster ipsDataItem_withPhoto">
                                                <li>
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/60659-wei%C3%9Fe-rose/"
                                                        data-ipshover=""
                                                        data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/60659-wei%C3%9Fe-rose/&amp;do=hovercard"
                                                        class="ipsUserPhoto ipsUserPhoto_tiny" title="Go to
Weiße Rose&#39;s profile">
                                                        <img src="https://kingdom-leaks.com/uploads/monthly_2019_10/csm_8065x_49b7a87273.thumb.jpg.63680ebda95e73ca449bda54c45ed895.jpg"
                                                            alt="Weiße Rose" itemprop="image" />
                                                    </a>
                                                </li>
                                                <li>
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/60659-wei%C3%9Fe-rose/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/60659-wei%C3%9Fe-rose/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.
php%253F%252Fforums%252Ftopic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F"
                                                        title="Go to Weiße Rose&#39;s profile"
                                                        class="ipsType_break"><span style="color:#2ecc71"><i
                                                                class="fa fa-usd"></i> Weiße Rose</span></a>
                                                </li>
                                                <li class="ipsType_light">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/17371-paradise-lost-believe-in-nothing-remixed-remastered-2018/&amp;do=getLastComment"
                                                        title="Go to last post" class="ipsType_blendLinks">
                                                        <time datetime="2018-07-01T12:54:29Z"
                                                            title="07/01/2018 12:54  PM" data-short="1 yr">July 1,
                                                            2018</time>
                                                    </a>
                                                </li>
                                            </ul>
                                        </li>
                                        <li class="ipsDataItem ipsDataItem_responsivePhoto   ">
                                            <div class="bim_tthumb_wrap" style="padding: 0px !important;">
                                                <div class="bim_tthumb_wrap">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/8756-paradise-lost-medusa-2017/"
                                                        data-ipstpopup="" data-ipstpopup-width="280"
                                                        data-ipstpopup-target="https://kingdom-leaks.com/index.php?/forums/topic/8756-paradise-lost-medusa-201
7/&amp;do=tthumbGetImage" alt="Paradise Lost - Medusa (2017)">
                                                        <div data-full="https://kingdom-leaks.com/uploads/monthly_2017_12/topicThumb_8756.jpg.d6187f132050e3a24fe038c893e70dab.jpg"
                                                            id="tthumb_8756" class="tthumb_standard lazyload" data-bg="https://kingdom-leaks.com/uploads/monthly_2017_1
2/topicThumb_8756.jpg.thumb.jpg.5a5fd461f7b1ebccd223f9660e1dd54d.jpg" style="width: 60px; height: 60px;"></div>
                                                    </a>
                                                    <!---->
                                                </div>
                                            </div>
                                            <div class="ipsDataItem_main">
                                                <h4 class="ipsDataItem_title ipsContained_container">
                                                    <span class="ipsType_break ipsContained">
                                                        <a href="https://kingdom-leaks.com/index.php?/forums/topic/8756-paradise-lost-medusa-2017/"
                                                            data-ipshover=""
                                                            data-ipshover-target="https://kingdom-leaks.com/index.php?/forums/topic/8756-paradise-lost-medusa-2017/&amp;preview=1"
                                                            data-ipshover-timeout="1.5">
                                                            Paradise Lost - Medusa (2017)
                                                        </a>
                                                    </span>
                                                </h4>
                                                <p class="ipsType_reset ipsType_medium ipsType_light">
                                                    By
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/899-why-nolin/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/899-why-nolin/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Ffor
ums%252Ftopic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Why Nolin?&#39;s profile"
                                                        class="ipsType_break"><span style="color:dimgrey"><i
                                                                class="fa fa-glass"></i> Why Nolin?</span></a>
                                                    <time datetime="2017-07-27T10:00:55Z" title="07/27/2017 10:00  AM"
                                                        data-short="2 yr">July 27, 2017</time>
                                                    in <a
                                                        href="https://kingdom-leaks.com/index.php?/forums/forum/11-metal/">Metal</a>
                                                </p>
                                                <ul class="ipsList_inline ipsClearfix ipsType_light">
                                                </ul>
                                            </div>
                                            <ul class="ipsDataItem_stats">
                                                <li>
                                                    <span class="ipsDataItem_stats_number">5</span>
                                                    <span class="ipsDataItem_stats_type"> replies</span>
                                                </li>
                                                <li>
                                                    <span class="ipsDataItem_stats_number">4,774</span>
                                                    <span class="ipsDataItem_stats_type"> views</span>
                                                </li>
                                            </ul>
                                            <ul class="ipsDataItem_lastPoster ipsDataItem_withPhoto">
                                                <li>
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/162-nimue/"
                                                        data-ipshover=""
                                                        data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/162-nimue/&amp;do=hovercard"
                                                        class="ipsUserPhoto ipsUserPhoto_tiny"
                                                        title="Go to Nimue&#39;s profile">
                                                        <img src="https://kingdom-leaks.com/uploads/monthly_2017_02/112e0e7ff71ef031ba304e6111c5797f.thumb.jpg.a02d27f4faad02d9f0444436d0c158e8.jpg"
                                                            alt="Nimue" itemprop="image" />
                                                    </a>
                                                </li>
                                                <li>
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/162-nimue/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/162-nimue/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Fforums%252F
topic%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Nimue&#39;s profile"
                                                        class="ipsType_break"><span style="color:dimgrey"><i
                                                                class="fa fa-glass"></i> Nimue</span></a>
                                                </li>
                                                <li class="ipsType_light">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/8756-paradise-lost-medusa-2017/&amp;do=getLastComment"
                                                        title="Go to last post" class="ipsType_blendLinks">
                                                        <time datetime="2017-07-30T13:58:51Z"
                                                            title="07/30/2017 01:58  PM" data-short="2 yr">July 30,
                                                            2017</time>
                                                    </a>
                                                </li>
                                            </ul>
                                        </li>
                                        <li class="ipsDataItem ipsDataItem_responsivePhoto   ">
                                            <div class="bim_tthumb_wrap" style="padding: 0px !important;">
                                                <div class="bim_tthumb_wrap">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/80-delta-heavy-paradise-lost-2016/"
                                                        data-ipstpopup="" data-ipstpopup-width="280"
                                                        data-ipstpopup-target="https://kingdom-leaks.com/index.php?/forums/topic/80-delta-heavy-paradise-lo
st-2016/&amp;do=tthumbGetImage" alt="Delta Heavy - Paradise Lost (2016)">
                                                        <div data-full="https://kingdom-leaks.com/uploads/monthly_2017_12/topicThumb_80.jpeg.cf4c15f499b2ed5e65a24e9fb57fab19.jpeg"
                                                            id="tthumb_80" class="tthumb_standard lazyload" data-bg="https://kingdom-leaks.com/uploads/monthly_2017_12/
topicThumb_80.jpeg.thumb.jpeg.2387e0ca12a582ab5af6a85f576e8820.jpeg" style="width: 60px; height: 60px;"></div>
                                                    </a>
                                                    <!---->
                                                </div>
                                            </div>
                                            <div class="ipsDataItem_main">
                                                <h4 class="ipsDataItem_title ipsContained_container">
                                                    <span class="ipsType_break ipsContained">
                                                        <a href="https://kingdom-leaks.com/index.php?/forums/topic/80-delta-heavy-paradise-lost-2016/"
                                                            data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/forums/topic/80-delta-heavy-paradise-lost-2016/&amp;prev
iew=1" data-ipshover-timeout="1.5">
                                                            Delta Heavy - Paradise Lost (2016)
                                                        </a>
                                                    </span>
                                                </h4>
                                                <p class="ipsType_reset ipsType_medium ipsType_light">
                                                    By
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/3-chief/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/3-chief/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Fforums%252Ftopi
c%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Chief&#39;s profile"
                                                        class="ipsType_break"><span style="color:#ff3300"><i
                                                                class="fa fa-shield"></i> Chief</span></a>
                                                    <time datetime="2016-03-05T19:02:43Z" title="03/05/2016 07:02  PM"
                                                        data-short="4 yr">March 5, 2016</time>
                                                    in <a
                                                        href="https://kingdom-leaks.com/index.php?/forums/forum/30-electronic/">Electronic</a>
                                                </p>
                                                <ul class="ipsList_inline ipsClearfix ipsType_light">
                                                </ul>
                                            </div>
                                            <ul class="ipsDataItem_stats">
                                                <li>
                                                    <span class="ipsDataItem_stats_number">4</span>
                                                    <span class="ipsDataItem_stats_type"> replies</span>
                                                </li>
                                                <li>
                                                    <span class="ipsDataItem_stats_number">1,352</span>
                                                    <span class="ipsDataItem_stats_type"> views</span>
                                                </li>
                                            </ul>
                                            <ul class="ipsDataItem_lastPoster ipsDataItem_withPhoto">
                                                <li>
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/3-chief/"
                                                        data-ipshover=""
                                                        data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/3-chief/&amp;do=hovercard"
                                                        class="ipsUserPhoto ipsUserPhoto_tiny"
                                                        title="Go to Chief&#39;s profile">
                                                        <img src="https://kingdom-leaks.com/uploads/monthly_2020_03/703E447E-010D-44A0-901D-3840063F368F.thumb.jpeg.753da89676f22573f1ca9f8e19bc3c08.jpeg"
                                                            alt="Chief" itemprop="image" />
                                                    </a>
                                                </li>
                                                <li>
                                                    <a href="https://kingdom-leaks.com/index.php?/profile/3-chief/"
                                                        data-ipshover="" data-ipshover-target="https://kingdom-leaks.com/index.php?/profile/3-chief/&amp;do=hovercard&amp;referrer=https%253A%252F%252Fkingdom-leaks.com%252Findex.php%253F%252Fforums%252Ftopi
c%252F37899-paradise-lost-obsidian-limited-edition-2020%252F" title="Go to Chief&#39;s profile"
                                                        class="ipsType_break"><span style="color:#ff3300"><i
                                                                class="fa fa-shield"></i> Chief</span></a>
                                                </li>
                                                <li class="ipsType_light">
                                                    <a href="https://kingdom-leaks.com/index.php?/forums/topic/80-delta-heavy-paradise-lost-2016/&amp;do=getLastComment"
                                                        title="Go to last post" class="ipsType_blendLinks">
                                                        <time datetime="2016-03-06T21:23:50Z"
                                                            title="03/06/2016 09:23  PM" data-short="4 yr">March 6,
                                                            2016</time>
                                                    </a>
                                                </li>
                                            </ul>
                                        </li>
                                    </ul>
                                    <span style="display:none" id="resultRaw">Paradise Lost - Believe In Nothing
                                        [Remixed &amp; Remastered] (2018)</span>
                                    <script>
                                        var resultRaw = document.getElementById('resultRaw').innerHTML;
                                        var artistRaw = document.getElementById('releaseTitle').innerHTML;
                                        var artistTitle = artistRaw.split(" - ")[0].trim();
                                        if (artistTitle == "Various Artists" || artistTitle == "VA") {
                                            document.getElementById("otherReleasesBy").style.display = "none";
                                        } else {
                                            var artistResult = resultRaw.split(" - ")[0].trim();
                                            var title = document.getElementById('relatedTitle');
                                            if (artistTitle.toLowerCase() != artistResult.toLowerCase()) {
                                                title.innerHTML = "Similar Releases";
                                            } else {
                                                title.innerHTML += artistTitle;
                                            }
                                            var moreResults = document.getElementById('moreResults');
                                            if (moreResults != null) {
                                                moreResults.href += artistTitle;
                                            }
                                        }
                                    </script>
                                </div>
                                <div id="similarTopicsBefore"></div>
                                <input type="hidden" name="csrfKey" value="538d6993577ac1a4f5e0fc418dc51e57" />
                            </form>
                        </div>
                        <a id="replyForm"></a>
                        <div data-role="replyArea"
                            class="cTopicPostArea ipsBox ipsBox_transparent ipsAreaBackground ipsPad cTopicPostArea_noSize ipsSpacer_top">
                            <div data-controller="core.global.core.login">
                                <input type="hidden" name="csrfKey" value="538d6993577ac1a4f5e0fc418dc51e57" />
                                <div class="ipsType_center ipsPad cGuestTeaser">
                                    <h2 class="ipsType_pageTitle">Create an account or sign in to comment</h2>
                                    <p class="ipsType_light ipsType_normal ipsType_reset">You need to be a member in
                                        order to leave a comment</p>
                                    <div class="ipsBox ipsPad ipsSpacer_top">
                                        <div class="ipsGrid ipsGrid_collapsePhone">
                                            <div class="ipsGrid_span6 cGuestTeaser_left">
                                                <h2 class="ipsType_sectionHead">Create an account</h2>
                                                <p class="ipsType_normal ipsType_reset ipsType_light ipsSpacer_bottom">
                                                    Sign up for a new account in our community. It&#39;s easy!</p>
                                                <a href="https://kingdom-leaks.com/index.php?/register/"
                                                    class="ipsButton ipsButton_primary ipsButton_small"
                                                    data-ipsdialog="" data-ipsdialog-size="narrow"
                                                    data-ipsdialog-title="Sign Up">
                                                    Register a new account</a>
                                            </div>
                                            <div class="ipsGrid_span6 cGuestTeaser_right">
                                                <h2 class="ipsType_sectionHead">Sign in</h2>
                                                <p class="ipsType_normal ipsType_reset ipsType_light ipsSpacer_bottom">
                                                    Already have an account? Sign in here.</p>
                                                <a href="https://kingdom-leaks.com/index.php?/login/&amp;ref=aHR0cHM6Ly9raW5nZG9tLWxlYWtzLmNvbS9pbmRleC5waHA/L2ZvcnVtcy90b3BpYy8zNzg5OS1wYXJhZGlzZS1sb3N0LW9ic2lkaWFuLWxpbWl0ZWQtZWRpdGlvbi0yMDIwLyNyZXBseUZvcm0="
                                                    data -ipsdialog="" data-ipsdialog-size="medium"
                                                    data-ipsdialog-title="Sign In Now"
                                                    class="ipsButton ipsButton_primary ipsButton_small">Sign In Now</a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="ipsResponsive_noFloat ipsResponsive_showPhone ipsResponsive_block ipsSpacer_top">
                            <div data-followapp="forums" data-followarea="topic" data-followid="37899"
                                data-controller="core.front.core.followButton">
                                <span
                                    class="ipsType_light ipsType_blendLinks ipsResponsive_hidePhone ipsResponsive_inline"><i
                                        class="fa fa-info-circle"></i> <a
                                        href="https://kingdom-leaks.com/index.php?/login/"
                                        title="Go to the sign in page">Sign in to follow this</a>  </span>
                                <div class="ipsFollow ipsPos_middle ipsButton ipsButton_link ipsButton_verySmall ipsButton_disabled"
                                    data-role="followButton">
                                    <span>Followers</span>
                                    <span class="ipsCommentCount">0</span>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="ipsGrid ipsGrid_collapsePhone ipsPager ipsClearfix ipsSpacer_top">
                        <div class="ipsGrid_span6 ipsType_left ipsPager_prev">
                            <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/"
                                title="Go to Currently Leaked" rel="up">
                                <span class="ipsPager_type">Go To Topic Listing</span>
                                <span class="ipsPager_title ipsType_light ipsTruncate ipsTruncate_line">Currently
                                    Leaked</span>
                            </a>
                        </div>
                    </div>
                </div>
                <nav class="ipsBreadcrumb ipsBreadcrumb_ ipsFaded_withHover">
                    <ul class="ipsList_inline ipsPos_right">
                        <li>
                            <a class="ipsType_light" href="https://status.kingdom-leaks.com/" target="_blank"><i
                                    class="fa fa-server" aria-hidden="true"></i> <span>KLStatus</span></a>
                        </li>
                    </ul>
                    <ul class="ipsList_inline ipsPos_right">
                        <li>
                            <a class="ipsType_light" href="/release-notes/"><i class="fa fa-sticky-note"
                                    aria-hidden="true"></i> <span>Release Notes </span></a>
                        </li>
                    </ul>
                    <ul class="ipsList_inline ipsPos_right">
                        <li>
                            <a class="ipsType_light" href="/new/"><i class="fa fa-newspaper-o" aria-hidden="true"></i>
                                <span>Latest Downloads </span></a>
                        </li>
                    </ul>
                    <ul data-role="breadcrumbList">
                        <li>
                            <a title="Home" href="https://kingdom-leaks.com/">
                                <span><i class="fa fa-home"></i> Home <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            <a href="https://kingdom-leaks.com/index.php?/forums/">
                                <span>Forums <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            <a href="https://kingdom-leaks.com/index.php?/forums/forum/108-leaks/">
                                <span>Leaks <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/">
                                <span>Currently Leaked <i class="fa fa-angle-right"></i></span>
                            </a>
                        </li>
                        <li>
                            Paradise Lost - Obsidian (Limited Edition) (2020)
                        </li>
                    </ul>
                </nav>
                <div style="text-align:center;margin-top:-22px;margin-bottom:8px;">
                    <ul class="veilon-social-icons ipsList_inline ipsPos_center"
                        style="margin:auto; margin-bottom: -13px;padding-left: 15px; max-width: 310px!important;">
                        <a title="Facebook" target="_blank" href="https://www.facebook.com/KingdomLeaks/
" data-ipstooltip=""><i class="fa fa-facebook ipsType_larger"></i></a>
                        <a title="Twitter" target="_blank" href="https://twitter.com/masterleakers
" data-ipstooltip=""><i class="fa fa-twitter ipsType_larger"></i></a>
                        <a title="Instagram" target="_blank" href="https://www.instagram.com/kingdomleaks/
" data-ipstooltip=""><i class="fa fa-instagram ipsType_larger"></i></a>
                        <a title="Spotify" target="_blank" href="https://open.spotify.com/user/p051x03x8hqzylgwgsg8ayt4r
" data-ipstooltip=""><i class="fa fa-spotify ipsType_larger"></i></a>
                        <a title="Apple Music" target="_blank" href="https://music.apple.com/profile/kingdomleaks
" data-ipstooltip=""><i class="fa fa-apple ipsType_larger"></i></a>
                        <a title="YouTube" target="_blank"
                            href="https://www.youtube.com/channel/UCxm0Vz6bzNFu5357QNQ5lyA" data-ipstooltip=""><i
                                class="fa fa-youtube-play ipsType_larger"></i></a>
                    </ul>
                    Copyright © 2013-2020 Kingdom Leaks.
                </div>
            </div>
        </div>
        <klstatus id="check_online"></klstatus>
        <a id="backtoTop" onclick="$(&#39;body&#39;).animatescroll();" title="Back To Top"><i
                class="fa fa-chevron-up"></i></a>
    </main>
    <div id="elMobileDrawer" class="ipsDrawer ipsHide">
        <a href="#" class="ipsDrawer_close" data-action="close"><span>×</span></a>
        <div class="ipsDrawer_menu">
            <div class="ipsDrawer_content">
                <div class="ipsSpacer_bottom ipsPad">
                    <ul class="ipsToolList ipsToolList_vertical">
                        <li>
                            <a href="//kingdom-leaks.com/index.php?/login/"
                                class="ipsButton ipsButton_light ipsButton_small ipsButton_fullWidth">Existing user?
                                Sign In</a>
                        </li>
                        <li>
                            <a href="//kingdom-leaks.com/index.php?/register/" id="elRegisterButton_mobile"
                                class="ipsButton ipsButton_small ipsButton_fullWidth ipsButton_important">Sign Up</a>
                        </li>
                        <!---->
                    </ul>
                </div>
                <ul class="ipsDrawer_list">
                    <li class="ipsDrawer_itemParent">
                        <h4 class="ipsDrawer_title"><a href="#">Browse</a></h4>
                        <ul class="ipsDrawer_list">
                            <li data-action="back"><a href="#">Back</a></li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/forums/">
                                    Forums
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/new/">
                                    Latest Downloads
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/">
                                    Currently Leaked
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/singles/">
                                    Latest Singles
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/trending/">
                                    Trending Leaks
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/hq/">
                                    Latest Quality Updates
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/pastleaders/">
                                    Leaderboard
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/online/&amp;filter=filter_loggedin">
                                    Online Users
                                </a>
                            </li>
                        </ul>
                    </li>
                    <li><a href="https://kingdom-leaks.com/index.php?/forms/2-submit-a-leak/">Submit a Leak</a></li>
                    <li class="ipsDrawer_itemParent">
                        <h4 class="ipsDrawer_title"><a href="#">About Us</a></h4>
                        <ul class="ipsDrawer_list">
                            <li data-action="back"><a href="#">Back</a></li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/staff/">
                                    Staff &amp; Trusted
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/rules/">
                                    Site Rules
                                </a>
                            </li>
                            <li>
                                <a
                                    href="https://kingdom-leaks.com/index.php?/forums/topic/28592-how-to-download-from-kingdom-leaks/">
                                    How to Download
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/ourpicks/">
                                    Our Picks
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/donations/">
                                    Donate
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/user-classes/">
                                    User Classes
                                </a>
                            </li>
                        </ul>
                    </li>
                    <li class="ipsDrawer_itemParent">
                        <h4 class="ipsDrawer_title"><a href="#">Misc</a></h4>
                        <ul class="ipsDrawer_list">
                            <li data-action="back"><a href="#">Back</a></li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/forums/forum/54-requests/">
                                    Request
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/forums/forum/54-requests/">
                                    Request
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/artwork/">
                                    Artwork Finder
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/chat/">
                                    Chat
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/release-notes/">
                                    Release Notes
                                </a>
                            </li>
                            <li>
                                <a href="https://kingdom-leaks.com/index.php?/feature-plan/">
                                    Feature Plan
                                </a>
                            </li>
                            <li>
                                <a href="https://status.kingdom-leaks.com/" target="_blank" rel="noopener">
                                    KLStatus
                                </a>
                            </li>
                            <li>
                                <a href="https://web.archive.org/save/kingdom-leaks.com" target="_blank" rel="noopener">
                                    #SaveKL
                                </a>
                            </li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
    </div>
    <script type="text/javascript">
        var ipsDebug = false;
        var CKEDITOR_BASEPATH = '//kingdom-leaks.com/applications/core/interface/ckeditor/ckeditor/';
        var ipsSettings = {
            cookie_path: "/",
            cookie_prefix: "ips4_",
            cookie_ssl: true,
            upload_imgURL: "",
            message_imgURL: "",
            notification_imgURL: "",
            baseURL: "//kingdom-leaks.com/",
            jsURL: "//kingdom-leaks.com/applications/core/interface/js/js.php",
            csrfKey: "538d6993577ac1a4f5e0fc418dc51e57",
            antiCache: "026d4e282a",
            disableNotificationSounds: false,
            useCompiledFiles: true,
            links_external: true,
            memberID: 0,
            analyticsProvider: "ga",
            viewProfiles: true,
            mapProvider: 'none',
            mapApiKey: '',
        };
    </script>
    <script type="text/javascript"
        src="//kingdom-leaks.com/applications/core/interface/howler/howler.core.min.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_global/root_library.js.82f4f90b27d21e35b8b7d03c4f455163.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_global/root_js_lang_1.js.17fafb3ba1491725362803a9c6860ba1.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_global/root_framework.js.4efd1545e23e99bfe56a41d77d85040b.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_core/global_global_core.js.cb16b9e175a9ea6c9efd678d0dee9fe6.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_core/plugins_plugins.js.af2699be13b16622405b13e43ce6bde3.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_bimchatbox/front_front_chatbox.js.aa92ba4f10ff65acc74145517679f663.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_global/root_front.js.0832846fe877cd151f630141287784dd.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_forums/front_front_topic.js.0dd532486e00ca7899c4640bca83bd9e.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_core/front_front_core.js.58f98d5f544446d26aa854675c872afb.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_featuredcontent/front_front_slider.js.b804ad6890806ce1530d1b7aaeaaba6d.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript"
        src="https://kingdom-leaks.com/uploads/javascript_global/root_map.js.bc7ba2852261e50e88d95d3e34f5aafd.js?v=026d4e282a"
        data-ips=""></script>
    <script type="text/javascript">
        ips.setSetting('date_format', jQuery.parseJSON('"mm\/dd\/yy"'));
        ips.setSetting('date_first_day', jQuery.parseJSON('0'));
        ips.setSetting('remote_image_proxy', jQuery.parseJSON('1'));
        ips.setSetting('ipb_url_filter_option', jQuery.parseJSON('"black"'));
        ips.setSetting('url_filter_any_action', jQuery.parseJSON('"allow"'));
        ips.setSetting('bypass_profanity', jQuery.parseJSON('0'));
        ips.setSetting('emoji_style', jQuery.parseJSON('"disabled"'));
        ips.setSetting('emoji_shortcodes', jQuery.parseJSON('"1"'));
        ips.setSetting('emoji_ascii', jQuery.parseJSON('"1"'));
        ips.setSetting('emoji_cache', jQuery.parseJSON('"1577940537"'));
        ips.setSetting('minimizeQuote_size', jQuery.parseJSON('5'));
        ips.setSetting('minimizeQuote_showFirstAppear', jQuery.parseJSON('0'));
        ips.setSetting('quickSearchDefault', jQuery.parseJSON('"forums_topic"'));
        ips.setSetting('canUseQuickSearch', jQuery.parseJSON('1'));
        ips.setSetting('quickSearchMinimum', jQuery.parseJSON('3'));
        ips.setSetting('quickSearchShowAdv', jQuery.parseJSON('true'));
        ips.setSetting('quickSearchIn', jQuery.parseJSON('"title"'));
    </script>
    <script type="application/ld+json">
{
"@context": "http://schema.org",
"@type": "DiscussionForumPosting",
"url": "https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/",
"discussionUrl": "https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/",
"name": "Paradise Lost - Obsidian (Limited Edition) (2020)",
"headline": "Paradise Lost - Obsidian (Limited Edition) (2020)",
"text": "Leaked\n\tRelease - May 15th, 2020\n\n\n\n\tGenre - Gothic Metal, Doom Metal\n\n\n\n\tQuality - MP3, 160 kbps CBR\n\n\n\n\t\u00a0\n\n\n\n\tTracklist:\n\n\n\n\t1. Darker Thoughts (05:46)\n\t2. Fall from Grace (05:42)\n\t3. Ghosts (04:35)\n\t4. The Dev
il Embraced (06:08)\n\t5. Forsaken (04:30)\n\t6. Serenity (04:46)\n\t7. Ending Days (04:36)\n\t8. Hope Dies Young (04:02)\n\t9. Ravenghast (05:30)\n\t10. Hear the Night (05:34)\n\t11. Defiler (04:45)\n\n\n\n\t\u00a0\n\n\n\n\tDownload\n\n\n\n\t\u00a0\n\n\n\n\tSupp
ort!\n\n\n\n\tFacebook / iTunes\n\n\n\n\t\u00a0\n\n\n\n\t\n\t\t\n\t\n\n\n\n\t\u00a0\n\n",
"dateCreated": "2020-05-04T17:00:56+0000",
"datePublished": "2020-05-04T17:00:56+0000",
"pageStart": 1,
"pageEnd": 1,
"image": "https://kingdom-leaks.com/uploads/monthly_2020_04/imageproxy.thumb.jpg.20df163213bc6f362a772b5aad24d2d5.jpg",
"author": {
"@type": "Person",
"name": "Summers",
"image": "https://kingdom-leaks.com/uploads/monthly_2020_04/imageproxy.thumb.jpg.20df163213bc6f362a772b5aad24d2d5.jpg",
"url": "https://kingdom-leaks.com/index.php?/profile/57459-summers/"
},
"interactionStatistic": [
{
"@type": "InteractionCounter",
"interactionType": "http://schema.org/ViewAction",
"userInteractionCount": 55
},
{
"@type": "InteractionCounter",
"interactionType": "http://schema.org/CommentAction",
"userInteractionCount": 1
},
{
"@type": "InteractionCounter",
"interactionType": "http://schema.org/FollowAction",
"userInteractionCount": 8
}
],
"comment": [
{
"@type": "Comment",
"url": "https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/\u0026do=findComment\u0026comment=229357",
"author": {
    "@type": "Person",
    "name": "Summers",
    "image": "https://kingdom-leaks.com/uploads/monthly_2020_04/imageproxy.thumb.jpg.20df163213bc6f362a772b5aad24d2d5.jpg",
    "url": "https://kingdom-leaks.com/index.php?/profile/57459-summers/"
},
"dateCreated": "2020-05-04T17:00:56+0000",
"text": "Leaked\n\tRelease - May 15th, 2020\n\n\n\n\tGenre - Gothic Metal, Doom Metal\n\n\n\n\tQuality - MP3, 160 kbps CBR\n\n\n\n\t\u00a0\n\n\n\n\tTracklist:\n\n\n\n\t1. Darker Thoughts (05:46)\n\t2. Fall from Grace (05:42)\n\t3. Ghosts (04:35)\n\t4.
The Devil Embraced (06:08)\n\t5. Forsaken (04:30)\n\t6. Serenity (04:46)\n\t7. Ending Days (04:36)\n\t8. Hope Dies Young (04:02)\n\t9. Ravenghast (05:30)\n\t10. Hear the Night (05:34)\n\t11. Defiler (04:45)\n\n\n\n\t\u00a0\n\n\n\n\tDownload\n\n\n\n\t\u00a0\n\n\n
\n\tSupport!\n\n\n\n\tFacebook / iTunes\n\n\n\n\t\u00a0\n\n\n\n\t\n\t\t\n\t\n\n\n\n\t\u00a0\n\n",
"mainEntityOfPage": "https://kingdom-leaks.com/index.php?/forums/topic/37899-paradise-lost-obsidian-limited-edition-2020/"
}
]
}
</script>
    <script type="application/ld+json">
{
"@context": "http://www.schema.org",
"@type": "WebSite",
"name": "Kingdom Leaks",
"url": "https://kingdom-leaks.com/",
"potentialAction": {
"type": "SearchAction",
"query-input": "required name=query",
"target": "https://kingdom-leaks.com/index.php?/search/\u0026q={query}"
},
"inLanguage": [
{
"@type": "Language",
"name": "English (USA)",
"alternateName": "en-US"
}
]
}
</script>
    <script type="application/ld+json">
{
"@context": "http://www.schema.org",
"@type": "Organization",
"name": "Kingdom Leaks",
"url": "https://kingdom-leaks.com/",
"logo": "https://kingdom-leaks.com/uploads/monthly_2020_02/KLll.png.577e903d7902e67c46afa23ca0fa6050.png",
"address": {
"@type": "PostalAddress",
"streetAddress": "",
"addressLocality": null,
"addressRegion": null,
"postalCode": null,
"addressCountry": null
}
}
</script>
    <script type="application/ld+json">
{
"@context": "http://schema.org",
"@type": "BreadcrumbList",
"itemListElement": [
{
"@type": "ListItem",
"position": 1,
"item": {
    "@id": "https://kingdom-leaks.com/index.php?/forums/",
    "name": "Forums"
}
},
{
"@type": "ListItem",
"position": 2,
"item": {
    "@id": "https://kingdom-leaks.com/index.php?/forums/forum/108-leaks/",
    "name": "Leaks"
}
},
{
"@type": "ListItem",
"position": 3,
"item": {
    "@id": "https://kingdom-leaks.com/index.php?/forums/forum/109-currently-leaked/",
    "name": "Currently Leaked"
}
}
]
}
</script>
    <!-- Google Analytics -->
    <script>
        (function (i, s, o, g, r, a, m) {
        i['GoogleAnalyticsObject'] = r; i[r] = i[r] || function () {
            (i[r].q = i[r].q || []).push(arguments)
        }, i[r].l = 1 * new Date(); a = s.createElement(o),
            m = s.getElementsByTagName(o)[0]; a.async = 1; a.src = g; m.parentNode.insertBefore(a, m)
        })(window, document, 'script', 'https://www.google-analytics.com/analytics.js', 'ga');
        ga('create', 'UA-88375452-1', 'auto');
        ga('send', 'pageview');
    </script>
    <!-- End Google Analytics -->
    <!--ipsQueryLog-->
    <!--ipsCachingLog-->
    <script type="text/javascript">
        $(document).ready(function () {
            $('.slider_1').bxSlider({
                wrapperClass: 'fcWrapper',
                mode: 'horizontal',
                slideMargin: 0,
                minSlides: 1,
                maxSlides: 8,
                slideWidth: 150,
                slideHeight: 150,
                adaptiveHeight: true,
                auto: false,
                autoDirection: 'next',
                autoControls: false,
                autoControlsCombine: false,
                controls: true,
                keyboardEnabled: true,
                autoHover: true,
                speed: 1000,
                ticker: false,
                tickerHover: true,
                pause: 5000,
                useCSS: false,
                touchEnabled: false,
                pager: false,
                onSliderLoad: function () {
                    setTimeout(function () {
                        $("#sliderWrapper_1").css({
                            "opacity": 1,
                            "height": "auto"
                        });
                    }, 300);
                },
            });
        });
    </script>
    <script>
        $(window).scroll(function () {
            var scroll = $(window).scrollTop();
            if (scroll >= 100) {
                $(".v-nav").addClass("v-sticky")
            } else {
                $(".v-nav").removeClass("v-sticky")
            }
        });
        //fix sticky nav mouseout issue
        ;; (function ($, _, undefined) {
            "use strict";
            ips.controller.register('core.front.core.navBar', {
                mouseOutScope: function () {
                    var self = this;
                    if (ips.utils.events.isTouchDevice()) {
                        return;
                    }
                    this._mouseOutTimer = setTimeout(function () {
                        self._makeDefaultActive();
                        //self.scope.find('[data-ipsMenu]').trigger('closeMenu');
                    }, 500);
                },
            });
        }(jQuery, _));
    </script>
    <script
        src="https://kingdom-leaks.com/uploads/set_resources_6/609fc90c74a47e3a2b9cd98294ff3404_nprogress.js"></script>
    <script>
        $('body').show();
        $('.version').text(NProgress.version);
        NProgress.start();
        setTimeout(function () { NProgress.done(); $('.fade').removeClass('out'); }, 1000);
    </script>
</body>

</html>
<nil>
`

func TestParseImage(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(site))
	if err != nil {
		t.Error(err)
		return
	}
	image, err := extractImage(doc)
	if err != nil {
		t.Error(err)
		return
	}

	expected := "https://kingdom-leaks.com/applications/core/interface/imageproxy/imageproxy.php?img=https://is5-ssl.mzstatic.com/image/thumb/Music114/v4/f4/f8/67/f4f8672e-b25b-5185-7b8d-27f54fb2af7c/source/600x600bb.jpg&key=a354a8baa5d145f970ae06fdea6bd929e5752c40a504258e8d750315ebfae55c"
	if !reflect.DeepEqual(expected, image) {
		t.Errorf("not equals:\n expected\t= %s\n actual\t\t= %s", expected, image)
	}
}
