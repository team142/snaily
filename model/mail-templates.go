package model

var WelcomeMailTemplate = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml">

<head>
    <!--[if gte mso 9]><xml><o:OfficeDocumentSettings><o:AllowPNG/><o:PixelsPerInch>96</o:PixelsPerInch></o:OfficeDocumentSettings></xml><![endif]-->
    <meta content="text/html; charset=utf-8" http-equiv="Content-Type" />
    <meta content="width=device-width" name="viewport" />
    <!--[if !mso]><!-->
    <meta content="IE=edge" http-equiv="X-UA-Compatible" />
    <!--<![endif]-->
    <title></title>
    <!--[if !mso]><!-->
    <link href="https://fonts.googleapis.com/css?family=Montserrat" rel="stylesheet" type="text/css" />
    <link href="https://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet" type="text/css" />
    <!--<![endif]-->
    <style type="text/css">
        body {
            margin: 0;
            padding: 0;
        }

        table,
        td,
        tr {
            vertical-align: top;
            border-collapse: collapse;
        }

        * {
            line-height: inherit;
        }

        a[x-apple-data-detectors=true] {
            color: inherit !important;
            text-decoration: none !important;
        }

        .ie-browser table {
            table-layout: fixed;
        }

        [owa] .img-container div,
        [owa] .img-container button {
            display: block !important;
        }

        [owa] .fullwidth button {
            width: 100% !important;
        }

        [owa] .block-grid .col {
            display: table-cell;
            float: none !important;
            vertical-align: top;
        }

        .ie-browser .block-grid,
        .ie-browser .num12,
        [owa] .num12,
        [owa] .block-grid {
            width: 700px !important;
        }

        .ie-browser .mixed-two-up .num4,
        [owa] .mixed-two-up .num4 {
            width: 232px !important;
        }

        .ie-browser .mixed-two-up .num8,
        [owa] .mixed-two-up .num8 {
            width: 464px !important;
        }

        .ie-browser .block-grid.two-up .col,
        [owa] .block-grid.two-up .col {
            width: 348px !important;
        }

        .ie-browser .block-grid.three-up .col,
        [owa] .block-grid.three-up .col {
            width: 348px !important;
        }

        .ie-browser .block-grid.four-up .col [owa] .block-grid.four-up .col {
            width: 174px !important;
        }

        .ie-browser .block-grid.five-up .col [owa] .block-grid.five-up .col {
            width: 140px !important;
        }

        .ie-browser .block-grid.six-up .col,
        [owa] .block-grid.six-up .col {
            width: 116px !important;
        }

        .ie-browser .block-grid.seven-up .col,
        [owa] .block-grid.seven-up .col {
            width: 100px !important;
        }

        .ie-browser .block-grid.eight-up .col,
        [owa] .block-grid.eight-up .col {
            width: 87px !important;
        }

        .ie-browser .block-grid.nine-up .col,
        [owa] .block-grid.nine-up .col {
            width: 77px !important;
        }

        .ie-browser .block-grid.ten-up .col,
        [owa] .block-grid.ten-up .col {
            width: 60px !important;
        }

        .ie-browser .block-grid.eleven-up .col,
        [owa] .block-grid.eleven-up .col {
            width: 54px !important;
        }

        .ie-browser .block-grid.twelve-up .col,
        [owa] .block-grid.twelve-up .col {
            width: 50px !important;
        }
    </style>
    <style id="media-query" type="text/css">
        @media only screen and (min-width: 720px) {
            .block-grid {
                width: 700px !important;
            }

            .block-grid .col {
                vertical-align: top;
            }

            .block-grid .col.num12 {
                width: 700px !important;
            }

            .block-grid.mixed-two-up .col.num3 {
                width: 174px !important;
            }

            .block-grid.mixed-two-up .col.num4 {
                width: 232px !important;
            }

            .block-grid.mixed-two-up .col.num8 {
                width: 464px !important;
            }

            .block-grid.mixed-two-up .col.num9 {
                width: 522px !important;
            }

            .block-grid.two-up .col {
                width: 350px !important;
            }

            .block-grid.three-up .col {
                width: 233px !important;
            }

            .block-grid.four-up .col {
                width: 175px !important;
            }

            .block-grid.five-up .col {
                width: 140px !important;
            }

            .block-grid.six-up .col {
                width: 116px !important;
            }

            .block-grid.seven-up .col {
                width: 100px !important;
            }

            .block-grid.eight-up .col {
                width: 87px !important;
            }

            .block-grid.nine-up .col {
                width: 77px !important;
            }

            .block-grid.ten-up .col {
                width: 70px !important;
            }

            .block-grid.eleven-up .col {
                width: 63px !important;
            }

            .block-grid.twelve-up .col {
                width: 58px !important;
            }
        }

        @media (max-width: 720px) {

            .block-grid,
            .col {
                min-width: 320px !important;
                max-width: 100% !important;
                display: block !important;
            }

            .block-grid {
                width: 100% !important;
            }

            .col {
                width: 100% !important;
            }

            .col>div {
                margin: 0 auto;
            }

            img.fullwidth,
            img.fullwidthOnMobile {
                max-width: 100% !important;
            }

            .no-stack .col {
                min-width: 0 !important;
                display: table-cell !important;
            }

            .no-stack.two-up .col {
                width: 50% !important;
            }

            .no-stack .col.num4 {
                width: 33% !important;
            }

            .no-stack .col.num8 {
                width: 66% !important;
            }

            .no-stack .col.num4 {
                width: 33% !important;
            }

            .no-stack .col.num3 {
                width: 25% !important;
            }

            .no-stack .col.num6 {
                width: 50% !important;
            }

            .no-stack .col.num9 {
                width: 75% !important;
            }

            .video-block {
                max-width: none !important;
            }

            .mobile_hide {
                min-height: 0px;
                max-height: 0px;
                max-width: 0px;
                display: none;
                overflow: hidden;
                font-size: 0px;
            }

            .desktop_hide {
                display: block !important;
                max-height: none !important;
            }
        }
    </style>
</head>

<body class="clean-body" style="margin: 0; padding: 0; -webkit-text-size-adjust: 100%; background-color: #FFFFFF;">
    <style id="media-query-bodytag" type="text/css">
        @media (max-width: 720px) {
            .block-grid {
                min-width: 320px!important;
                max-width: 100%!important;
                width: 100%!important;
                display: block!important;
            }
            .col {
                min-width: 320px!important;
                max-width: 100%!important;
                width: 100%!important;
                display: block!important;
            }
            .col>div {
                margin: 0 auto;
            }
            img.fullwidth {
                max-width: 100%!important;
                height: auto!important;
            }
            img.fullwidthOnMobile {
                max-width: 100%!important;
                height: auto!important;
            }
            .no-stack .col {
                min-width: 0!important;
                display: table-cell!important;
            }
            .no-stack.two-up .col {
                width: 50%!important;
            }
            .no-stack.mixed-two-up .col.num4 {
                width: 33%!important;
            }
            .no-stack.mixed-two-up .col.num8 {
                width: 66%!important;
            }
            .no-stack.three-up .col.num4 {
                width: 33%!important
            }
            .no-stack.four-up .col.num3 {
                width: 25%!important
            }
        }
    </style>
    <!--[if IE]><div class="ie-browser"><![endif]-->
    <table bgcolor="#FFFFFF" cellpadding="0" cellspacing="0" class="nl-container" role="presentation" style="table-layout: fixed; vertical-align: top; min-width: 320px; Margin: 0 auto; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF; width: 100%;"
        valign="top" width="100%">
        <tbody>
            <tr style="vertical-align: top;" valign="top">
                <td style="word-break: break-word; vertical-align: top; border-collapse: collapse;" valign="top">
                    <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td align="center" style="background-color:#FFFFFF"><![endif]-->
                    <div style="background-color:#FFFFFF;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#FFFFFF;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:transparent;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:0px; padding-bottom:0px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:0px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <table border="0" cellpadding="0" cellspacing="0" class="divider" role="presentation" style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;"
                                                valign="top" width="100%">
                                                <tbody>
                                                    <tr style="vertical-align: top;" valign="top">
                                                        <td class="divider_inner" style="word-break: break-word; vertical-align: top; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; padding-top: 0px; padding-right: 0px; padding-bottom: 0px; padding-left: 0px; border-collapse: collapse;"
                                                            valign="top">
                                                            <table align="center" border="0" cellpadding="0" cellspacing="0" class="divider_content" height="30" role="presentation"
                                                                style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; border-top: 0px solid transparent; height: 30px;"
                                                                valign="top" width="100%">
                                                                <tbody>
                                                                    <tr style="vertical-align: top;" valign="top">
                                                                        <td height="30" style="word-break: break-word; vertical-align: top; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; border-collapse: collapse;"
                                                                            valign="top">
                                                                            <span></span>
                                                                        </td>
                                                                    </tr>
                                                                </tbody>
                                                            </table>
                                                        </td>
                                                    </tr>
                                                </tbody>
                                            </table>
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-color:transparent;">
                        <div class="block-grid two-up no-stack" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:transparent;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="350" style="background-color:transparent;width:350px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 20px; padding-left: 20px; padding-top:10px; padding-bottom:10px;"><![endif]-->
                                <div class="col num6" style="min-width: 320px; max-width: 350px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:10px; padding-bottom:10px; padding-right: 20px; padding-left: 20px;">
                                            <!--<![endif]-->
                                            <div></div>
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td><td align="center" width="350" style="background-color:transparent;width:350px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px;"><![endif]-->
                                <div class="col num6" style="min-width: 320px; max-width: 350px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 20px; padding-bottom: 10px; font-family: Arial, sans-serif"><![endif]-->
                                            <div style="color:#E3E3E3;font-family:'Open Sans', Helvetica, Arial, sans-serif;line-height:120%;padding-top:20px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                                                <div style="font-size: 12px; line-height: 14px; font-family: 'Open Sans', Helvetica, Arial, sans-serif; color: #E3E3E3;">
                                                    <p style="font-size: 14px; line-height: 14px; text-align: right; margin: 0;">
                                                        <span style="font-size: 12px;"></span>
                                                    </p>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-image:url('images/Bg_text_ani.gif');background-position:top center;background-repeat:no-repeat;background-color:#FFFFFF;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-image:url('images/Bg_text_ani.gif');background-position:top center;background-repeat:no-repeat;background-color:#FFFFFF;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:transparent;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:40px; padding-bottom:5px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:40px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <div></div>
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-image:url('images/bg_wave_1.png');background-position:top center;background-repeat:repeat;background-color:#F4F4F4;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-image:url('images/bg_wave_1.png');background-position:top center;background-repeat:repeat;background-color:#F4F4F4;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:transparent;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:0px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <table border="0" cellpadding="0" cellspacing="0" class="divider" role="presentation" style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;"
                                                valign="top" width="100%">
                                                <tbody>
                                                    <tr style="vertical-align: top;" valign="top">
                                                        <td class="divider_inner" style="word-break: break-word; vertical-align: top; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; padding-top: 10px; padding-right: 10px; padding-bottom: 10px; padding-left: 10px; border-collapse: collapse;"
                                                            valign="top">
                                                            <table align="center" border="0" cellpadding="0" cellspacing="0" class="divider_content" height="70" role="presentation"
                                                                style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; border-top: 0px solid transparent; height: 70px;"
                                                                valign="top" width="100%">
                                                                <tbody>
                                                                    <tr style="vertical-align: top;" valign="top">
                                                                        <td height="70" style="word-break: break-word; vertical-align: top; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; border-collapse: collapse;"
                                                                            valign="top">
                                                                            <span></span>
                                                                        </td>
                                                                    </tr>
                                                                </tbody>
                                                            </table>
                                                        </td>
                                                    </tr>
                                                </tbody>
                                            </table>
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-color:#F4F4F4;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F4F4F4;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:transparent;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:5px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <img width="25%" src="http://icons.iconarchive.com/icons/google/noto-emoji-animals-nature/1024/22303-snail-icon.png" />
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 30px; padding-left: 30px; padding-top: 10px; padding-bottom: 0px; font-family: 'Trebuchet MS', Tahoma, sans-serif"><![endif]-->
                                            <div style="color:#555555;font-family:'Montserrat', 'Trebuchet MS', 'Lucida Grande', 'Lucida Sans Unicode', 'Lucida Sans', Tahoma, sans-serif;line-height:120%;padding-top:10px;padding-right:30px;padding-bottom:0px;padding-left:30px;">
                                                <div style="font-size: 12px; line-height: 14px; font-family: 'Montserrat', 'Trebuchet MS', 'Lucida Grande', 'Lucida Sans Unicode', 'Lucida Sans', Tahoma, sans-serif; color: #555555;">
                                                    <p style="font-size: 14px; line-height: 16px; margin: 0;">
                                                        <strong>
                                                            <span style="font-size: 46px; line-height: 55px;">Welcome,
                                                                <span style="color: #3d3bee; line-height: 55px; font-size: 46px;">XXX</span>!</span>
                                                        </strong>
                                                    </p>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 30px; padding-left: 30px; padding-top: 15px; padding-bottom: 5px; font-family: Arial, sans-serif"><![endif]-->
                                            <div style="color:#555555;font-family:'Open Sans', Helvetica, Arial, sans-serif;line-height:150%;padding-top:15px;padding-right:30px;padding-bottom:5px;padding-left:30px;">
                                                <div style="font-size: 12px; line-height: 18px; font-family: 'Open Sans', Helvetica, Arial, sans-serif; color: #555555;">
                                                    <p style="font-size: 12px; line-height: 18px; margin: 0;">
                                                        <strong>
                                                            <span style="font-size: 20px; line-height: 30px;">DependMap helps teams and companies find out where things get slow
                                                                or stuck. You're recieving this email because you have been
                                                                signed up or because someone... is waiting... for you! </span>
                                                        </strong>
                                                    </p>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 30px; padding-left: 30px; padding-top: 15px; padding-bottom: 20px; font-family: Arial, sans-serif"><![endif]-->
                                            <div style="color:#7C7C7C;font-family:'Open Sans', Helvetica, Arial, sans-serif;line-height:150%;padding-top:15px;padding-right:30px;padding-bottom:20px;padding-left:30px;">
                                                <div style="font-size: 12px; line-height: 18px; font-family: 'Open Sans', Helvetica, Arial, sans-serif; color: #7C7C7C;">
                                                    <p style="font-size: 12px; line-height: 24px; margin: 0;">
                                                        <span style="font-size: 16px;">Don't freak out. This app is not designed to find out who the problem
                                                            is but rather to help identify where things get stuck. That way
                                                            the company can change how it operates and improve it's processes
                                                            to ensure things people like you wait less.
                                                            <span
                                                                style="text-decoration: underline; font-size: 16px; line-height: 24px;">It's not about targeting you. It's about targeting the things
                                                                you wait for!</span>
                                                        </span>
                                                    </p>
                                                    <p style="font-size: 12px; line-height: 18px; margin: 0;"> </p>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-color:#F4F4F4;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #FFFFFF;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:#FFFFFF;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F4F4F4;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:#FFFFFF"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:#FFFFFF;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 5px; padding-left: 5px; padding-top:5px; padding-bottom:5px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 5px; padding-left: 5px;">
                                            <!--<![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 35px; padding-top: 25px; padding-bottom: 0px; font-family: 'Trebuchet MS', Tahoma, sans-serif"><![endif]-->
                                            <div style="color:#444444;font-family:'Montserrat', 'Trebuchet MS', 'Lucida Grande', 'Lucida Sans Unicode', 'Lucida Sans', Tahoma, sans-serif;line-height:120%;padding-top:25px;padding-right:10px;padding-bottom:0px;padding-left:35px;">
                                                <div style="font-size: 12px; line-height: 14px; font-family: 'Montserrat', 'Trebuchet MS', 'Lucida Grande', 'Lucida Sans Unicode', 'Lucida Sans', Tahoma, sans-serif; color: #444444;">
                                                    <p style="font-size: 14px; line-height: 16px; margin: 0;">
                                                        <span style="background-color: #93f5ed; font-size: 14px; line-height: 16px;">
                                                            <span style="font-size: 24px; line-height: 28px; background-color: #93f5ed;">
                                                                <span style="line-height: 28px; font-size: 24px; background-color: #93f5ed;"> List of features: </span>
                                                            </span>
                                                        </span>
                                                    </p>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 20px; padding-left: 30px; padding-top: 15px; padding-bottom: 30px; font-family: Arial, sans-serif"><![endif]-->
                                            <div style="color:#555555;font-family:'Open Sans', Helvetica, Arial, sans-serif;line-height:180%;padding-top:15px;padding-right:20px;padding-bottom:30px;padding-left:30px;">
                                                <div style="line-height: 21px; font-family: 'Open Sans', Helvetica, Arial, sans-serif; font-size: 12px; color: #555555;">
                                                    <ul style="font-size: 12px;">
                                                        <li style="line-height: 21px; font-size: 12px;">
                                                            <span style="font-size: 16px; line-height: 28px;">Create a new
                                                                <em>Snaily</em> when you're waiting for someone.</span>
                                                        </li>
                                                        <li style="font-size: 14px; line-height: 25px;">
                                                            <span style="font-size: 16px; line-height: 28px;">CC new@dependmap.com when you're raising a blockage in your area
                                                                and new snaily is created. </span>
                                                        </li>
                                                        <li style="font-size: 14px; line-height: 25px;">
                                                            <span style="font-size: 16px; line-height: 28px;">Loads of reports. I mean all of them! Almost.</span>
                                                        </li>
                                                        <li style="font-size: 14px; line-height: 25px;">
                                                            <span style="font-size: 16px; line-height: 28px;">More on the way.</span>
                                                        </li>
                                                    </ul>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-color:#F4F4F4;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F4F4F4;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:transparent;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:25px; padding-bottom:60px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:25px; padding-bottom:60px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <div align="center" class="button-container" style="padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                                                <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="border-spacing: 0; border-collapse: collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;"><tr><td style="padding-top: 10px; padding-right: 10px; padding-bottom: 10px; padding-left: 10px" align="center"><v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="" style="height:54pt; width:228.75pt; v-text-anchor:middle;" arcsize="13%" stroke="false" fillcolor="#3D3BEE"><w:anchorlock/><v:textbox inset="0,0,0,0"><center style="color:#ffffff; font-family:Arial, sans-serif; font-size:26px"><![endif]-->
                                                <div style="text-decoration:none;display:inline-block;color:#ffffff;background-color:#3D3BEE;border-radius:9px;-webkit-border-radius:9px;-moz-border-radius:9px;width:auto; width:auto;;border-top:1px solid #3D3BEE;border-right:1px solid #3D3BEE;border-bottom:1px solid #3D3BEE;border-left:1px solid #3D3BEE;padding-top:10px;padding-bottom:10px;font-family:'Open Sans', Helvetica, Arial, sans-serif;text-align:center;mso-border-alt:none;word-break:keep-all;">
                                                    <span style="padding-left:45px;padding-right:45px;font-size:26px;display:inline-block;">
                                                        <a href="https://dependmap.com">
                                                            <span style="font-size: 16px; line-height: 32px;">
                                                                <span style="font-size: 26px; line-height: 52px; color:white">
                                                                    <strong>Dependmap.com</strong>
                                                                </span>
                                                            </span>
                                                        </a>
                                                    </span>
                                                </div>
                                                <!--[if mso]></center></v:textbox></v:roundrect></td></tr></table><![endif]-->
                                            </div>
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-image:url('images/bg_wave_2.png');background-position:top center;background-repeat:repeat;background-color:#F4F4F4;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: transparent;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-image:url('images/bg_wave_2.png');background-position:top center;background-repeat:repeat;background-color:#F4F4F4;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:transparent"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:transparent;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:5px; padding-bottom:0px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:0px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <table border="0" cellpadding="0" cellspacing="0" class="divider" role="presentation" style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;"
                                                valign="top" width="100%">
                                                <tbody>
                                                    <tr style="vertical-align: top;" valign="top">
                                                        <td class="divider_inner" style="word-break: break-word; vertical-align: top; min-width: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; padding-top: 10px; padding-right: 10px; padding-bottom: 10px; padding-left: 10px; border-collapse: collapse;"
                                                            valign="top">
                                                            <table align="center" border="0" cellpadding="0" cellspacing="0" class="divider_content" height="70" role="presentation"
                                                                style="table-layout: fixed; vertical-align: top; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; border-top: 0px solid transparent; height: 70px;"
                                                                valign="top" width="100%">
                                                                <tbody>
                                                                    <tr style="vertical-align: top;" valign="top">
                                                                        <td height="70" style="word-break: break-word; vertical-align: top; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; border-collapse: collapse;"
                                                                            valign="top">
                                                                            <span></span>
                                                                        </td>
                                                                    </tr>
                                                                </tbody>
                                                            </table>
                                                        </td>
                                                    </tr>
                                                </tbody>
                                            </table>
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <div style="background-color:#FFFFFF;">
                        <div class="block-grid" style="Margin: 0 auto; min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; background-color: #FFFFFF;;">
                            <div style="border-collapse: collapse;display: table;width: 100%;background-color:#FFFFFF;">
                                <!--[if (mso)|(IE)]><table width="100%" cellpadding="0" cellspacing="0" border="0" style="background-color:#FFFFFF;"><tr><td align="center"><table cellpadding="0" cellspacing="0" border="0" style="width:700px"><tr class="layout-full-width" style="background-color:#FFFFFF"><![endif]-->
                                <!--[if (mso)|(IE)]><td align="center" width="700" style="background-color:#FFFFFF;width:700px; border-top: 0px solid transparent; border-left: 0px solid transparent; border-bottom: 0px solid transparent; border-right: 0px solid transparent;" valign="top"><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 0px; padding-left: 0px; padding-top:15px; padding-bottom:35px;"><![endif]-->
                                <div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top;;">
                                    <div style="width:100% !important;">
                                        <!--[if (!mso)&(!IE)]><!-->
                                        <div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:15px; padding-bottom:35px; padding-right: 0px; padding-left: 0px;">
                                            <!--<![endif]-->
                                            <!--[if mso]><table width="100%" cellpadding="0" cellspacing="0" border="0"><tr><td style="padding-right: 10px; padding-left: 10px; padding-top: 10px; padding-bottom: 10px; font-family: Arial, sans-serif"><![endif]-->
                                            <div style="color:#838383;font-family:'Open Sans', Helvetica, Arial, sans-serif;line-height:150%;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
                                                <div style="font-size: 12px; line-height: 18px; font-family: 'Open Sans', Helvetica, Arial, sans-serif; color: #838383;">
                                                    <p style="font-size: 14px; line-height: 21px; text-align: center; margin: 0;">
                                                        <span style="color: #000000; font-size: 14px; line-height: 21px;">
                                                            <strong>DependMap.com</strong>
                                                        </span>, helping you find your bottleneck since 2018.</p>
                                                    <p style="font-size: 14px; line-height: 21px; text-align: center; margin: 0;">173 Bush Hill Manor, South Africa   </p>
                                                </div>
                                            </div>
                                            <!--[if mso]></td></tr></table><![endif]-->
                                            <!--[if (!mso)&(!IE)]><!-->
                                        </div>
                                        <!--<![endif]-->
                                    </div>
                                </div>
                                <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                                <!--[if (mso)|(IE)]></td></tr></table></td></tr></table><![endif]-->
                            </div>
                        </div>
                    </div>
                    <!--[if (mso)|(IE)]></td></tr></table><![endif]-->
                </td>
            </tr>
        </tbody>
    </table>
    <!--[if (IE)]></div><![endif]-->
</body>

</html>


`